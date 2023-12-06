package handlers

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

// MainHandler
func MainHandler(content embed.FS) http.HandlerFunc {
	log.Print("MainHandler called")
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(content, "templates/main.html", "templates/monsterForm.html", "templates/monster.html", "templates/monsterTable.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(w, "main", map[string]interface{}{
			"Title": "Dungeons & Dragons Monster Generator",
		})
		if err != nil {
			log.Printf("Template execution error: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
