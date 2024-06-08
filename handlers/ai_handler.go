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
				prompt = "Generate a Random Homebrew Dnd monster. You can choose the CR, the type, and the design. Give me a detailed sheet for the monster."
			} else {
				name := r.FormValue("name")
				cr := r.FormValue("cr")
				monsterType := r.FormValue("type")
				monsterInfo := r.FormValue("monsterinfo")
				prompt = "Create a DnD monster with the following details:\n"
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
				"max_tokens":  800,
				"temperature": 0.7,
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
				filename := "generated_monster.txt"
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
