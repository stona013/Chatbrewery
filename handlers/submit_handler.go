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

// submitHandler verarbeitet die Formulardaten
func SubmitHandler(content embed.FS, chars *[]model.Character, Monsters *[]model.Monster, filename string) http.HandlerFunc {
	log.Print("SubmitHandler called")
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Formulardaten parsen
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Monster-Objekt erstellen
		filename := r.FormValue("filename")

		// Charakter-Objekt erstellen oder aktualisieren
		mu.Lock()
		defer mu.Unlock()

		char := model.GetOrCreateCharacter(filename, *chars)
		char.Monster = append(char.Monster, *Monsters...)

		// Charakterdaten in JSON umwandeln
		charJSON, err := json.Marshal(char)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Monster hinzugefügt. Anzahl der Monster jetzt: %d\n", len(*Monsters))
		// JSON-Daten in die Datei schreiben
		err = model.WriteToFile(filename, charJSON)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Dateiinhalt lesen
		fileContent, err := os.ReadFile(filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Datei zum Download anbieten
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
		w.Header().Set("Content-Type", "application/json")
		w.Write(fileContent)
	}
}
