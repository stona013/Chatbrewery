package handlers

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

// ContactHandler handles the contact page request.
// It takes the content embed.FS as a parameter and returns an http.HandlerFunc.
// The returned http.HandlerFunc renders the contact page using the provided templates.
func ContactHandler(content embed.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("ContactHandler called")

		// Parse the templates
		tmpl, err := template.ParseFS(content, "templates/base.html", "templates/header.html", "templates/main.html", "templates/footer.html", "templates/about.html", "templates/contact.html")
		if err != nil {
			log.Printf("Template parsing error: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute the contact template
		err = tmpl.ExecuteTemplate(w, "contact", map[string]interface{}{
			"Title": "Dungeons & Dragons Monster Generator",
		})
		if err != nil {
			log.Printf("Template execution error: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
