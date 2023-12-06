package handlers

import (
	"ddServer/model"
	"embed"
	"html/template"
	"log"
	"net/http"
)

// MonsterTableHandler returns a http.HandlerFunc that handles requests to display a table of monsters.
func MonsterTableHandler(content embed.FS, monsters *[]model.Monster) http.HandlerFunc {
	log.Print("MonsterTableHandler called")

	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Handling request for monster table")

		// Parse the template files
		tmpl, err := template.ParseFS(content, "templates/base.html", "templates/header.html", "templates/main.html", "templates/footer.html", "templates/monsterTable.html", "templates/monster.html")
		if err != nil {
			log.Printf("Template parsing error: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute the template and pass the necessary data
		err = tmpl.ExecuteTemplate(w, "monsterTable", map[string]interface{}{
			"Title":    "Dungeons & Dragons Monster Generator",
			"Monsters": *monsters,
		})
		if err != nil {
			log.Printf("Template execution error: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
