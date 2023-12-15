package handlers

import (
	"ddServer/model"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	var (
		acrobatics     = ""
		animalHandling = ""
		arcana         = ""
		athletics      = ""
		deception      = ""
		history        = ""
		insight        = ""
		intimidation   = ""
		investigation  = ""
		medicine       = ""
		nature         = ""
		performance    = ""
		perception     = ""
		persuasion     = ""
		sleightOfHand  = ""
		religion       = ""
		stealth        = ""
		survival       = ""
	)
	if r.FormValue("checkAcrobatics") == "on" {
		acrobatics = r.FormValue("acrobatics")
	}
	if r.FormValue("checkAnimalHandling") == "on" {
		animalHandling = r.FormValue("animalHandling")
	}
	if r.FormValue("checkArcana") == "on" {
		arcana = r.FormValue("arcana")
	}
	if r.FormValue("checkAthletics") == "on" {
		athletics = r.FormValue("athletics")
	}
	if r.FormValue("checkDeception") == "on" {
		deception = r.FormValue("deception")
	}
	if r.FormValue("checkHistory") == "on" {
		history = r.FormValue("history")
	}
	if r.FormValue("checkInsight") == "on" {
		insight = r.FormValue("insight")
	}
	if r.FormValue("checkIntimidation") == "on" {
		intimidation = r.FormValue("intimidation")
	}
	if r.FormValue("checkInvestigation") == "on" {
		investigation = r.FormValue("investigation")
	}
	if r.FormValue("checkMedicine") == "on" {
		medicine = r.FormValue("medicine")
	}
	if r.FormValue("checkNature") == "on" {
		nature = r.FormValue("nature")
	}
	if r.FormValue("checkPerformance") == "on" {
		performance = r.FormValue("performance")
	}
	if r.FormValue("checkPerception") == "on" {
		perception = r.FormValue("perception")
	}
	if r.FormValue("checkPersuasion") == "on" {
		persuasion = r.FormValue("persuasion")
	}
	if r.FormValue("checkSleightOfHand") == "on" {
		sleightOfHand = r.FormValue("sleightOfHand")
	}
	if r.FormValue("checkStealth") == "on" {
		stealth = r.FormValue("stealth")
	}
	if r.FormValue("checkSurvival") == "on" {
		survival = r.FormValue("survival")
	}
	if r.FormValue("checkReligion") == "on" {
		religion = r.FormValue("religion")
	}
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
			Dex: r.FormValue("saveDex"),
			Con: r.FormValue("saveCon"),
			Wis: r.FormValue("saveWis"),
			Str: r.FormValue("saveStr"),
			Cha: r.FormValue("saveCha"),
			Int: r.FormValue("saveInt"),
		},
		Skill: model.Skill{
			Perception:     perception,
			Stealth:        stealth,
			Acrobatics:     acrobatics,
			AnimalHandling: animalHandling,
			Arcana:         arcana,
			Athletics:      athletics,
			Deception:      deception,
			History:        history,
			Insight:        insight,
			Intimidation:   intimidation,
			Investigation:  investigation,
			Medicine:       medicine,
			Nature:         nature,
			Performance:    performance,
			Persuasion:     persuasion,
			SleightOfHand:  sleightOfHand,
			Survival:       survival,
			Religion:       religion,
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
