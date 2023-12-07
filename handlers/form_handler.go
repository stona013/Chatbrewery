package handlers

import (
	"ddServer/model"
	"embed"
	"html/template"
	"log"
	"net/http"
)

// FormHandler returns an http.HandlerFunc that handles form submissions.
// It takes the content embed.FS, a pointer to a slice of model.Monster,
// and a filename string as parameters.
// The function parses the template files from the content FS,
// executes the template with the provided data, and renders it as a response.
func FormHandler(content embed.FS, monsters *[]model.Monster) http.HandlerFunc {
	log.Print("FormHandler called")

	// Lock the mutex to ensure exclusive access to the monsters slice.
	mu.Lock()
	defer mu.Unlock()

	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("FormHandler handler called")

		// Parse the template files.
		templateFiles := []string{
			"templates/base.html",
			"templates/header.html",
			"templates/main.html",
			"templates/footer.html",
			"templates/monsterForm.html",
			"templates/monster.html",
			"templates/monsterTable.html",
		}
		tmpl, err := template.ParseFS(content, templateFiles...)
		if err != nil {
			log.Printf("Template parsing error: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute the template and render the response.
		data := map[string]interface{}{
			"Title":    "Dungeons & Dragons Monster Generator",
			"Monsters": *monsters,
		}
		err = tmpl.ExecuteTemplate(w, "base", data)
		if err != nil {
			log.Printf("Template execution error: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Printf("Template rendered with %d Monsters\n", len(*monsters))
	}
}
