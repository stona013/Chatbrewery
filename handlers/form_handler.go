package handlers

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

func FormHandler(content embed.FS, filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(content, "templates/base.html", "templates/header.html", "templates/main.html", "templates/footer.html", "templates/monsterForm.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(w, "base", map[string]interface{}{
			"Title": "Dungeons & Dragons Monster Generator",
		})
		if err != nil {
			log.Printf("Template execution error: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
