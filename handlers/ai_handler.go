package handlers

import (
	"bytes"
	"ddServer/model"
	"embed"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// Add a function to handle file download
func downloadFile(w http.ResponseWriter, r *http.Request, filePath string) {
	w.Header().Set("Content-Disposition", "attachment; filename="+filePath)
	w.Header().Set("Content-Type", "text/plain")
	http.ServeFile(w, r, filePath)
}

func AIHandler(content embed.FS, monsters *[]model.Monster) http.HandlerFunc {
	log.Print("AIHandler called")

	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("AIHandler request received")

		if r.Method == http.MethodPost {
			apiKey := r.FormValue("apikey")
			generationType := r.FormValue("generationType")

			var prompt string
			if generationType == "random" {
				prompt = `"Create a random detailed DnD monster for the plutonium importer tool 5etools for the inport on Foundry Vtt with the following Json structure but input the stats and features from the monster you have builded:" 
json structure:
{  "_meta":{"_dateLastModifiedHash":"66658f76","sources":[{"json":"chatbrewery","abbreviation":"MG","version":"unknown","authors":["Krzysztof"],"convertedBy":["Krzysztof"]}],"dateAdded":1717931894,"dateLastModified":1717931894},"monster":[{"save":{"dex":"1","con":"1","wis":"1","cha":"1","str":"1","int":"1"},"skill":{"stealth":"-3","acrobatics":"-3","animalHandling":"-3","arcana":"-3","athletics":"-3","deception":"-3","history":"-3","insight":"-3","intimidation":"-3","investigation":"-3","medicine":"-3","nature":"-3","perception":"-3","performance":"-3","persuasion":"-3","sleightOfHand":"-3","survival":"-3","religion":"-3"},"hp":{"formula":"1","average":1},"source":"1","cr":"1","type":"beast","name":"test","vulnerable":["1"],"conditionImmune":["1"],"resist":["1"],"immune":["1"],"trait":[{"name":"1","entries":["1"]}],"ac":[{"from":["1"],"ac":1}],"senses":["1"],"languages":["1"],"size":["H"],"action":[{"name":"Claw","entries":["{@atk mw} {@hit 7} to hit, reach 5 ft., one target. {@h}17 ({@damage 2d12 + 5}) bludgeoning damage. On a hit, the target must make a DC 16 Strength saving throw or be knocked prone."]}],"speed":{"walk":1,"burrow":1,"climb":1,"fly":1,"swim":1},"str":1,"dex":1,"con":1,"int":1,"wis":1,"cha":1}]}" "
you can choose the CR the Type and the Name for the Monster and just response with the Json structure and nothing else not even a comand form you and also dont put backtick att the beginning and the end! 
`
			} else {
				name := r.FormValue("name")
				cr := r.FormValue("cr")
				monsterType := r.FormValue("type")
				monsterInfo := r.FormValue("monsterinfo")
				prompt += `"Create a detailed DnD monster for the plutonium importer tool from 5etools for the inport on Foundry Vtt with the following Json structure just response with the Json structure and nothing else not even a comand form you and also dont put backtick att the beginning and the end!!:
json structure :
{  "_meta":{"_dateLastModifiedHash":"66658f76","sources":[{"json":"Malgorgon","abbreviation":"MG","version":"unknown","authors":["Krzysztof"],"convertedBy":["Krzysztof"]}],"dateAdded":1717931894,"dateLastModified":1717931894},"monster":[{"save":{"dex":"1","con":"1","wis":"1","cha":"1","str":"1","int":"1"},"skill":{"stealth":"-3","acrobatics":"-3","animalHandling":"-3","arcana":"-3","athletics":"-3","deception":"-3","history":"-3","insight":"-3","intimidation":"-3","investigation":"-3","medicine":"-3","nature":"-3","perception":"-3","performance":"-3","persuasion":"-3","sleightOfHand":"-3","survival":"-3","religion":"-3"},"hp":{"formula":"1","average":1},"source":"1","cr":"1","type":"beast","name":"test","vulnerable":["1"],"conditionImmune":["1"],"resist":["1"],"immune":["1"],"trait":[{"name":"1","entries":["1"]}],"ac":[{"from":["1"],"ac":1}],"senses":["1"],"languages":["1"],"size":["H"],"action":[{"name":"Claw","entries":["{@atk mw} {@hit 7} to hit, reach 5 ft., one target. {@h}17 ({@damage 2d12 + 5}) bludgeoning damage. On a hit, the target must make a DC 16 Strength saving throw or be knocked prone."]}],"speed":{"walk":1,"burrow":1,"climb":1,"fly":1,"swim":1},"str":1,"dex":1,"con":1,"int":1,"wis":1,"cha":1}]}
` + "\n"

				if name != "" {
					prompt += "Name: " + name + "\n"
				}
				if cr != "" {
					prompt += "CR: " + cr + "\n"
				}
				if monsterType != "" {
					prompt += "Type: " + monsterType + "\n"
				}
				if monsterInfo != "" {
					prompt += "Details: " + monsterInfo + "\n"
				}
			}

			data := map[string]interface{}{
				"model": "gpt-4-turbo",
				"messages": []map[string]string{
					{"role": "user", "content": prompt},
				},
				"max_tokens":  1500,
				"temperature": 0.5,
			}

			requestData, err := json.Marshal(data)
			if err != nil {
				log.Printf("Error marshalling JSON: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestData))
			if err != nil {
				log.Printf("Error creating request: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+apiKey)

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Printf("Error sending request: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			defer resp.Body.Close()

			responseBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Error reading response: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			var aiResponse OpenAIResponse
			if err := json.Unmarshal(responseBody, &aiResponse); err != nil {
				log.Printf("Error unmarshalling JSON: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if len(aiResponse.Choices) > 0 {
				messageContent := aiResponse.Choices[0].Message.Content

				// Save the generated monster to a TXT file
				filename := "generated_monster.json"
				err := ioutil.WriteFile(filename, []byte(messageContent), 0644)
				if err != nil {
					log.Printf("Error writing to file: %v\n", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}

				// Trigger the file download
				downloadFile(w, r, filename)
				return
			}
		}

		tmplFiles := []string{"templates/base.html", "templates/header.html", "templates/main.html", "templates/footer.html", "templates/ai.html"}
		tmpl, err := template.ParseFS(content, tmplFiles...)
		if err != nil {
			log.Printf("Template parsing error: %v\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data := map[string]interface{}{
			"Title": "Dungeons & Dragons Monster Generator",
		}
		err = tmpl.ExecuteTemplate(w, "ai", data)
		if err != nil {
			log.Printf("Template execution error: %v\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
