package handlers

import (
	"ddServer/model"
	"embed"
	"html/template"
	"log"
	"net/http"
)

// MainHandler handles the main HTTP request.
// It returns an http.HandlerFunc that renders the main page
// with the provided content and monsters.
func MainHandler(content embed.FS, monsters *[]model.Monster) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("MainHandler called")

		// Parse the templates from the embedded file system
		tmpl, err := template.ParseFS(content, "templates/main.html", "templates/monsterForm.html", "templates/monster.html", "templates/monsterTable.html", "templates/base.html", "templates/skills.html")
		if err != nil {
			log.Printf("Template parsing error: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Lock the mutex to ensure exclusive access to the monsters slice
		mu.Lock()
		defer mu.Unlock()

		// Execute the main template with the provided data
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
