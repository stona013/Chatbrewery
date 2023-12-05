package handlers

import (
	"ddServer/model"
	"embed"
	"html/template"
	"log"
	"net/http"
)

func FormHandler(content embed.FS, monsters *[]model.Monster, filename string) http.HandlerFunc {
	log.Print("FormHandler called")
	mu.Lock()
	defer mu.Unlock()
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(content, "templates/base.html", "templates/header.html", "templates/main.html", "templates/footer.html", "templates/monsterForm.html", "templates/monster.html", "templates/monsterTable.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(w, "base", map[string]interface{}{
			"Title":    "Dungeons & Dragons Monster Generator",
			"Monsters": *monsters,
		})
		if err != nil {
			log.Printf("Template execution error: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Printf("Template mit %d Monstern gerendert\n", len(*monsters))
	}
}
