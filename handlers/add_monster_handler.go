package handlers

import (
	"ddServer/model"
	"log"
	"net/http"
	"strconv"
)

func AddMonster(Monsters *[]model.Monster) http.HandlerFunc {
	log.Print("AddMonster called")
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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
			DamageRes: []string{r.FormValue("damageRes")},
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
		*Monsters = append(*Monsters, monster)
	}
}

// parseInt konvertiert einen String zu einem Integer und gibt 0 zurück, wenn die Konvertierung fehlschlägt
func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
