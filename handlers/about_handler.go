package handlers

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

// AboutHandler returns an http.HandlerFunc that handles requests to the /about endpoint.
// It renders the about.html template and passes in the title "Dungeons & Dragons Monster Generator".
func AboutHandler(content embed.FS) http.HandlerFunc {
	log.Print("AboutHandler called")

	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("AboutHandler request received")

		// Parse the template files
		tmplFiles := []string{"templates/base.html", "templates/header.html", "templates/main.html", "templates/footer.html", "templates/about.html"}
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
		err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			log.Printf("Template execution error: %v\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
