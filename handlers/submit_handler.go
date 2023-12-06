package handlers

import (
	"ddServer/model"
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var mu sync.Mutex

// SubmitHandler processes the form data.
func SubmitHandler(content embed.FS, chars *[]model.Character, Monsters *[]model.Monster, filename string) http.HandlerFunc {
	log.Print("SubmitHandler called")
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("SubmitHandler called")
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse form data.
		err := r.ParseForm()
		if err != nil {
			log.Printf("Error parsing form data: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create monster object.
		filename := r.FormValue("filename")

		// Create or update character object.
		mu.Lock()
		defer mu.Unlock()

		char := model.GetOrCreateCharacter(filename, *chars)
		char.Monster = append(char.Monster, *Monsters...)

		// Convert character data to JSON.
		charJSON, err := json.Marshal(char)
		if err != nil {
			log.Printf("Error marshalling character data to JSON: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write JSON data to file.
		err = model.WriteToFile(filename, charJSON)
		if err != nil {
			log.Printf("Error writing JSON data to file: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Read file contents.
		fileContent, err := os.ReadFile(filename)
		if err != nil {
			log.Printf("Error reading file contents: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Offer file for download.
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
		w.Header().Set("Content-Type", "application/json")
		w.Write(fileContent)
	}
}
