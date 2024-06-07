package handlers

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

// AboutHandler returns an http.HandlerFunc that handles requests to the /about endpoint.
// It renders the about.html template and passes in the title "Dungeons & Dragons Monster Generator".
func AIHandler(content embed.FS) http.HandlerFunc {
	log.Print("AIHandler called")

	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("AIHandler request received")

		// Parse the template files
		tmplFiles := []string{"templates/base.html", "templates/header.html", "templates/main.html", "templates/footer.html", "templates/ai.html"}
		tmpl, err := template.ParseFS(content, tmplFiles...)
		if err != nil {
			log.Printf("Template parsing error: %v\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Execute the template with the provided data
		data := map[string]interface{}{
			"Title": "Dungeons & Dragons Monster Generator",
		}
		err = tmpl.ExecuteTemplate(w, "ai", data)
		if err != nil {
			log.Printf("Template execution error: %v\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
