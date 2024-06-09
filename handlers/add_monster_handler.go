package handlers

import (
	"ddServer/model"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
			http.Error(w, err.Error(), http.StatusNoContent)
			return
		}

		// Create a new monster with the form data
		monster := parseMonster(r)

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

// parseMonster parses the Monster from monsterForm.html and return it.
func parseMonster(r *http.Request) model.Monster {
	return model.Monster{
		Name:      r.FormValue("name"),
		Source:    r.FormValue("source"),
		Size:      []string{r.FormValue("size")},
		Type:      strings.ToLower(r.FormValue("type")),
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
			Walk:   parseInt(r.FormValue("walk")),
			Burrow: parseInt(r.FormValue("burrow")),
			Fly:    parseInt(r.FormValue("fly")),
			Swim:   parseInt(r.FormValue("swim")),
			Climb:  parseInt(r.FormValue("climb")),
		},
		Str: parseInt(r.FormValue("str")),
		Dex: parseInt(r.FormValue("dex")),
		Con: parseInt(r.FormValue("con")),
		Int: parseInt(r.FormValue("int")),
		Wis: parseInt(r.FormValue("wis")),
		Cha: parseInt(r.FormValue("cha")),
		Save: model.Save{
			Dex: checkCheckbox("savedex", r),
			Con: checkCheckbox("savecon", r),
			Wis: checkCheckbox("savewis", r),
			Str: checkCheckbox("savestr", r),
			Cha: checkCheckbox("savecha", r),
			Int: checkCheckbox("saveint", r),
		},
		Skill: model.Skill{
			Perception:     checkCheckbox("perception", r),
			Stealth:        checkCheckbox("stealth", r),
			Acrobatics:     checkCheckbox("acrobatics", r),
			AnimalHandling: checkCheckbox("animalhandling", r),
			Arcana:         checkCheckbox("arcana", r),
			Athletics:      checkCheckbox("athletics", r),
			Deception:      checkCheckbox("deception", r),
			History:        checkCheckbox("history", r),
			Insight:        checkCheckbox("insight", r),
			Intimidation:   checkCheckbox("intimidation", r),
			Investigation:  checkCheckbox("investigation", r),
			Medicine:       checkCheckbox("medicine", r),
			Nature:         checkCheckbox("nature", r),
			Performance:    checkCheckbox("performance", r),
			Persuasion:     checkCheckbox("persuasion", r),
			SleightOfHand:  checkCheckbox("sleightofhand", r),
			Survival:       checkCheckbox("survival", r),
			Religion:       checkCheckbox("religion", r),
		},
		Resist:          []string{r.FormValue("resist")},
		ConditionImmune: []string{r.FormValue("conditionImmune")},
		Immune:          []string{r.FormValue("immune")},
		Vulnerable:      []string{r.FormValue("vulnerable")},
		Senses:          []string{r.FormValue("senses")},
		Languages:       []string{r.FormValue("languages")},
		CR:              r.FormValue("cr"),
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
}

func checkCheckbox(field string, r *http.Request) string {
	if r.FormValue(fmt.Sprintf("check%v", cases.Caser(cases.Title(language.Und)).String(field))) == "on" {
		return r.FormValue(field)
	}
	return ""
}
