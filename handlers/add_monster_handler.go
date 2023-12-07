package handlers

import (
	"ddServer/model"
	"log"
	"net/http"
	"strconv"
)

// AddMonster is a http.HandlerFunc that adds a new monster to the Monsters slice.
// It expects a POST request with form data containing the details of the monster.
// The monster is then appended to the Monsters slice and a redirect response is sent.
func AddMonster(Monsters *[]model.Monster) http.HandlerFunc {
	log.Print("AddMonster called")
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST
		if r.Method != http.MethodPost {
			log.Print("Method not allowed")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			log.Printf("Error parsing form data: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create a new monster with the form data
		monster := model.Monster{
			Name:      r.FormValue("name"),
			Source:    r.FormValue("source"),
			Size:      []string{r.FormValue("size")},
			Type:      r.FormValue("type"),
			Alignment: []string{r.FormValue("alignment")},
			AC: []model.AC{
				{
					AC:   parseInt(r.FormValue("ac")),
					From: []string{r.FormValue("acFrom")},
				},
			},
			HP: model.HP{
				Average: parseInt(r.FormValue("hpAverage")),
				Formula: r.FormValue("hpFormula"),
			},
			Speed: model.Speed{
				Walk: parseInt(r.FormValue("speed")),
			},
			Str: parseInt(r.FormValue("str")),
			Dex: parseInt(r.FormValue("dex")),
			Con: parseInt(r.FormValue("con")),
			Int: parseInt(r.FormValue("int")),
			Wis: parseInt(r.FormValue("wis")),
			Cha: parseInt(r.FormValue("cha")),
			Save: model.Save{
				Dex: r.FormValue("saveDex"),
				Con: r.FormValue("saveCon"),
				Wis: r.FormValue("saveWis"),
			},
			Skill: model.Skill{
				Perception: r.FormValue("perception"),
				Stealth:    r.FormValue("stealth"),
			},
			Resist:    []string{r.FormValue("resist")},
			Senses:    []string{r.FormValue("senses")},
			Languages: []string{r.FormValue("languages")},
			CR:        r.FormValue("cr"),
			Traits: []model.Trait{
				{
					Name:    r.FormValue("traitName"),
					Entries: []string{r.FormValue("traitEntry")},
				},
			},
			Actions: []model.Action{
				{
					Name:    r.FormValue("actionName"),
					Entries: []string{r.FormValue("actionEntry")},
				},
			},
		}

		// Lock the Monsters slice, append the monster, and unlock the slice
		mu.Lock()
		defer mu.Unlock()
		*Monsters = append(*Monsters, monster)

		// Log the number of monsters and redirect to the monster table
		log.Printf("Monster added. Number of monsters now: %d\n", len(*Monsters))
		http.Redirect(w, r, "/monsterTable", http.StatusFound)
	}
}

// parseInt converts a string to an integer and returns 0 if the conversion fails
func parseInt(s string) int {
	// Add logging statement to print the input string
	log.Println("Input string:", s)

	// Atoi is used to convert the string to an integer
	i, err := strconv.Atoi(s)
	// If there is an error in the conversion, return 0 and log the error
	if err != nil {
		log.Println("Conversion error:", err)
		return 0
	}
	// Log the converted integer
	log.Println("Converted integer:", i)

	// Return the converted integer
	return i
}
