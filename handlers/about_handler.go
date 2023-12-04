package handlers

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

func AboutHandler(content embed.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(content, "templates/base.html", "templates/header.html", "templates/main.html", "templates/footer.html", "templates/about.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(w, "about", map[string]interface{}{
			"Title": "Dungeons & Dragons Monster Generator",
		})
		if err != nil {
			log.Printf("Template execution error: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
