package handlers

import (
	"ddServer/model"
	"embed"
	"html/template"
	"log"
	"net/http"
)

// MainHandler
func MainHandler(content embed.FS, monsters *[]model.Monster) http.HandlerFunc {
	log.Print("MainHandler called")
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(content, "templates/main.html", "templates/monsterForm.html", "templates/monster.html", "templates/monsterTable.html", "templates/base.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		mu.Lock()
		defer mu.Unlock()

		err = tmpl.ExecuteTemplate(w, "main", map[string]interface{}{
			"Title":    "Dungeons & Dragons Monster Generator",
			"Monsters": *monsters,
		})
		if err != nil {
			log.Printf("Template execution error: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
