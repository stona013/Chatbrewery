package handlers

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// SkillCalculationHandler ist ein http.HandlerFunc, der von htmx getriggert wird,
// wenn der Benutzer Einträge in bestimmten Feldern macht, und dann die Skill-Felder befüllt.
func SkillCalculationHandler(content embed.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("SkillCalculationHandler called")

		// Überprüfen Sie, ob die Anfrage eine POST-Anfrage ist.
		if r.Method != http.MethodPost {
			log.Print("Method not allowed")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse Formulardaten.
		err := r.ParseForm()
		if err != nil {
			log.Printf("Error parsing form data: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmplFiles := []string{"templates/base.html", "templates/header.html", "templates/skills.html", "templates/main.html", "templates/footer.html", "templates/about.html"}
		tmpl, err := template.ParseFS(content, tmplFiles...)
		if err != nil {
			log.Printf("Template parsing error: %v\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		str := parseFieldValue(r.FormValue("str"))
		dex := parseFieldValue(r.FormValue("dex"))
		int := parseFieldValue(r.FormValue("int"))
		cha := parseFieldValue(r.FormValue("cha"))
		wis := parseFieldValue(r.FormValue("wis"))
		cr := parseFieldValue(r.FormValue("cr"))
		crBonus := calcBonus(cr)

		skillValues := map[string]string{
			"acrobatics":     strconv.Itoa(calcAbilityScore(dex) + crBonus),
			"animalHandling": strconv.Itoa(calcAbilityScore(wis) + crBonus),
			"arcana":         strconv.Itoa(calcAbilityScore(int) + crBonus),
			"athletics":      strconv.Itoa(calcAbilityScore(str) + crBonus),
			"deception":      strconv.Itoa(calcAbilityScore(cha) + crBonus),
			"history":        strconv.Itoa(calcAbilityScore(int) + crBonus),
			"insight":        strconv.Itoa(calcAbilityScore(wis) + crBonus),
			"intimidation":   strconv.Itoa(calcAbilityScore(cha) + crBonus),
			"investigation":  strconv.Itoa(calcAbilityScore(int) + crBonus),
			"medicine":       strconv.Itoa(calcAbilityScore(wis) + crBonus),
			"nature":         strconv.Itoa(calcAbilityScore(int) + crBonus),
			"perception":     strconv.Itoa(calcAbilityScore(wis) + crBonus),
			"performance":    strconv.Itoa(calcAbilityScore(cha) + crBonus),
			"persuasion":     strconv.Itoa(calcAbilityScore(cha) + crBonus),
			"religion":       strconv.Itoa(calcAbilityScore(int) + crBonus),
			"sleightOfHand":  strconv.Itoa(calcAbilityScore(dex) + crBonus),
			"stealth":        strconv.Itoa(calcAbilityScore(dex) + crBonus),
			"survival":       strconv.Itoa(calcAbilityScore(wis) + crBonus),
		}

		err = tmpl.ExecuteTemplate(w, "skills", skillValues)
		if err != nil {
			log.Printf("Template execution error: %v\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func calcBonus(cr int) int {
	if cr >= 0 && cr < 5 {
		return 2
	} else if cr >= 5 && cr < 9 {
		return 3
	} else if cr >= 9 && cr < 14 {
		return 4
	} else if cr >= 14 && cr < 18 {
		return 5
	} else if cr >= 18 && cr < 21 {
		return 6
	} else if cr >= 21 && cr < 25 {
		return 7
	} else if cr >= 25 && cr < 28 {
		return 8
	} else if cr >= 28 && cr < 31 {
		return 9
	} else {
		return 0
	}
}

func calcAbilityScore(val int) int {
	if val < 2 {
		return -5
	} else if val >= 2 && val < 4 {
		return -4
	} else if val >= 4 && val < 6 {
		return -3
	} else if val >= 6 && val < 8 {
		return -2
	} else if val >= 8 && val < 10 {
		return -1
	} else if val >= 10 && val < 12 {
		return 0
	} else if val >= 12 && val < 14 {
		return 1
	} else if val >= 14 && val < 16 {
		return 2
	} else if val >= 16 && val < 18 {
		return 3
	} else if val >= 18 && val < 20 {
		return 4
	} else if val >= 20 && val < 22 {
		return 5
	} else if val >= 22 && val < 24 {
		return 6
	} else if val >= 24 && val < 26 {
		return 7
	} else if val >= 26 && val < 28 {
		return 8
	} else if val >= 28 && val < 30 {
		return 9
	} else {
		return 10
	}
}

func parseFieldValue(value string) int {
	val, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Error converting field value to integer: %v", err)
		return 0
	}
	return val
}
