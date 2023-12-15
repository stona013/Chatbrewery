package handlers

import (
	"ddServer/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func LoadFileHandler(monsters *[]model.Monster) http.HandlerFunc {
	log.Print("LoadFileHandler called")
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20) // 10 MB limit

		// Get the file from the request
		file, _, err := r.FormFile("uploadFile")
		if err != nil {
			http.Error(w, "Error retrieving file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Parse the file content
		decoder := json.NewDecoder(file)
		var loadedChars model.Character
		err = decoder.Decode(&loadedChars)
		if err != nil {
			http.Error(w, "Error decoding file content", http.StatusInternalServerError)
			return
		}

		// Lock the Monsters slice and append the loaded monsters, then unlock the slice
		mu.Lock()
		defer mu.Unlock()

		// Assuming 'loadedChars' contains an array of Monster objects
		for _, monster := range loadedChars.Monster {
			*monsters = append(*monsters, monster)
		}

		fmt.Printf("%v\n", monsters)
		// Send a success response
		http.Redirect(w, r, "/monsterTable", http.StatusTemporaryRedirect)
	}
}
