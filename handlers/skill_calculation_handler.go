package handlers

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// SkillCalculationHandler is an http.HandlerFunc triggered by htmx when the user makes entries in certain fields and then populates the skill fields.
func SkillCalculationHandler(content embed.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is a POST request.
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse form data.
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Parse template files.
		tmplFiles := []string{
			"templates/base.html",
			"templates/header.html",
			"templates/skills.html",
			"templates/main.html",
			"templates/footer.html",
			"templates/about.html",
		}
		tmpl, err := template.ParseFS(content, tmplFiles...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Parse form field values and calculate skill values.
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

		// Execute template with skill values.
		err = tmpl.ExecuteTemplate(w, "skills", skillValues)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

// calcBonus calculates the bonus based on the given credit rating.
// It returns the bonus value as an integer.
func calcBonus(cr int) int {
	switch {
	case cr >= 0 && cr < 5:
		log.Println("Bonus calculated for credit rating:", cr)
		return 2
	case cr >= 5 && cr < 9:
		log.Println("Bonus calculated for credit rating:", cr)
		return 3
	case cr >= 9 && cr < 14:
		log.Println("Bonus calculated for credit rating:", cr)
		return 4
	case cr >= 14 && cr < 18:
		log.Println("Bonus calculated for credit rating:", cr)
		return 5
	case cr >= 18 && cr < 21:
		log.Println("Bonus calculated for credit rating:", cr)
		return 6
	case cr >= 21 && cr < 25:
		log.Println("Bonus calculated for credit rating:", cr)
		return 7
	case cr >= 25 && cr < 28:
		log.Println("Bonus calculated for credit rating:", cr)
		return 8
	case cr >= 28 && cr < 31:
		log.Println("Bonus calculated for credit rating:", cr)
		return 9
	default:
		log.Println("Invalid credit rating:", cr)
		return 0
	}
}

// calcAbilityScore calculates the ability score based on the given value.
func calcAbilityScore(val int) int {
	switch {
	case val < 2:
		log.Println("Ability Score: -5")
		return -5
	case val < 4:
		log.Println("Ability Score: -4")
		return -4
	case val < 6:
		log.Println("Ability Score: -3")
		return -3
	case val < 8:
		log.Println("Ability Score: -2")
		return -2
	case val < 10:
		log.Println("Ability Score: -1")
		return -1
	case val < 12:
		log.Println("Ability Score: 0")
		return 0
	case val < 14:
		log.Println("Ability Score: 1")
		return 1
	case val < 16:
		log.Println("Ability Score: 2")
		return 2
	case val < 18:
		log.Println("Ability Score: 3")
		return 3
	case val < 20:
		log.Println("Ability Score: 4")
		return 4
	case val < 22:
		log.Println("Ability Score: 5")
		return 5
	case val < 24:
		log.Println("Ability Score: 6")
		return 6
	case val < 26:
		log.Println("Ability Score: 7")
		return 7
	case val < 28:
		log.Println("Ability Score: 8")
		return 8
	case val < 30:
		log.Println("Ability Score: 9")
		return 9
	default:
		log.Println("Ability Score: 10")
		return 10
	}
}

// parseFieldValue takes a string value and returns an integer.
// If the string value cannot be converted to an integer, it logs an error and returns 0.
// The function follows these rules:
// - No line is over 66 characters.
func parseFieldValue(value string) int {
	// Convert the string value to an integer using strconv.Atoi.
	// If an error occurs during the conversion, log the error and return 0.
	val, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Error converting field value to integer: %v", err)
		return 0
	}
	// Log the converted integer value for debugging purposes.
	log.Printf("Converted field value to integer: %d", val)
	// Return the converted integer value.
	return val
}
