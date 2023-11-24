package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

// Monster struct für die Daten des Monsters
type Monster struct {
	Name      string   `json:"name"`
	Source    string   `json:"source"`
	Size      []string `json:"size"`
	Type      string   `json:"type"`
	Alignment []string `json:"alignment"`
	AC        []AC     `json:"ac"`
	HP        HP       `json:"hp"`
	Speed     Speed    `json:"speed"`
	Save      Save     `json:"save"`
	Skill     Skill    `json:"skill"`
	DamageRes []string `json:"damageResistances"`
	Senses    []string `json:"senses"`
	Languages []string `json:"languages"`
	CR        string   `json:"cr"`
	Traits    []Trait  `json:"trait"`
	Actions   []Action `json:"action"`
	Str       int      `json:"str"`
	Dex       int      `json:"dex"`
	Con       int      `json:"con"`
	Int       int      `json:"int"`
	Wis       int      `json:"wis"`
	Cha       int      `json:"cha"`
}

type AC struct {
	AC   int      `json:"ac"`
	From []string `json:"from"`
}

type HP struct {
	Average int    `json:"average"`
	Formula string `json:"formula"`
}

type Speed struct {
	Walk int `json:"walk"`
}

type Save struct {
	Dex string `json:"dex"`
	Con string `json:"con"`
	Wis string `json:"wis"`
}

type Skill struct {
	Perception string `json:"perception"`
	Stealth    string `json:"stealth"`
}

type Trait struct {
	Name    string   `json:"name"`
	Entries []string `json:"entries"`
}

type Action struct {
	Name    string   `json:"name"`
	Entries []string `json:"entries"`
}

// Character struct für die Daten des Charakters
type Character struct {
	Meta    Meta      `json:"_meta"`
	Monster []Monster `json:"monster"`
}

// Meta struct für Meta-Informationen
type Meta struct {
	Sources              []Source `json:"sources"`
	DateAdded            int64    `json:"dateAdded"`
	DateLastModified     int64    `json:"dateLastModified"`
	DateLastModifiedHash string   `json:"_dateLastModifiedHash"`
}

type Source struct {
	Json         string   `json:"json"`
	Abbreviation string   `json:"abbreviation"`
	Authors      []string `json:"authors"`
	ConvertedBy  []string `json:"convertedBy"`
	Version      string   `json:"version"`
}

var (
	mu    sync.Mutex
	chars []Character
	//go:embed forms.html
	page embed.FS
)

func main() {
	filename := ""

	http.HandleFunc("/", formHandler(filename))
	http.HandleFunc("/submit", submitHandler(filename))

	fmt.Println("Server gestartet, erreichbar unter http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// formHandler zeigt das Formular an
func formHandler(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(page, "forms.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	}
}

// submitHandler verarbeitet die Formulardaten
func submitHandler(filename string) http.HandlerFunc {
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
		monster := Monster{
			Name:      r.FormValue("name"),
			Source:    r.FormValue("source"),
			Size:      []string{r.FormValue("size")},
			Type:      r.FormValue("type"),
			Alignment: []string{r.FormValue("alignment")},
			AC: []AC{
				{
					AC:   parseInt(r.FormValue("ac")),
					From: []string{r.FormValue("acFrom")},
				},
			},
			HP: HP{
				Average: parseInt(r.FormValue("hpAverage")),
				Formula: r.FormValue("hpFormula"),
			},
			Speed: Speed{
				Walk: parseInt(r.FormValue("speed")),
			},
			Str: parseInt(r.FormValue("str")),
			Dex: parseInt(r.FormValue("dex")),
			Con: parseInt(r.FormValue("con")),
			Int: parseInt(r.FormValue("int")),
			Wis: parseInt(r.FormValue("wis")),
			Cha: parseInt(r.FormValue("cha")),
			Save: Save{
				Dex: r.FormValue("saveDex"),
				Con: r.FormValue("saveCon"),
				Wis: r.FormValue("saveWis"),
			},
			Skill: Skill{
				Perception: r.FormValue("perception"),
				Stealth:    r.FormValue("stealth"),
			},
			DamageRes: []string{r.FormValue("damageRes")},
			Senses:    []string{r.FormValue("senses")},
			Languages: []string{r.FormValue("languages")},
			CR:        r.FormValue("cr"),
			Traits: []Trait{
				{
					Name:    r.FormValue("traitName"),
					Entries: []string{r.FormValue("traitEntry")},
				},
			},
			Actions: []Action{
				{
					Name:    r.FormValue("actionName"),
					Entries: []string{r.FormValue("actionEntry")},
				},
			},
		}

		// Charakter-Objekt erstellen oder aktualisieren
		mu.Lock()
		defer mu.Unlock()

		char := getOrCreateCharacter(filename)
		char.Monster = append(char.Monster, monster)

		// Charakterdaten in JSON umwandeln
		charJSON, err := json.Marshal(char)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// JSON-Daten in die Datei schreiben
		err = writeToFile(filename, charJSON)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Erfolgsmeldung anzeigen
		fmt.Fprintf(w, "Monsterdaten erfolgreich gespeichert in %s: %s", filename, charJSON)
	}
}

// writeToFile schreibt Daten in eine Datei
func writeToFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// getOrCreateCharacter gibt das aktuelle Charakterobjekt zurück oder erstellt ein neues
func getOrCreateCharacter(filename string) Character {
	for _, char := range chars {
		if char.Meta.DateLastModified == 0 {
			// Ein leeres Charakterobjekt wurde gefunden
			return char
		}
	}

	// Erstelle ein neues Charakterobjekt
	now := time.Now().Unix()
	newChar := Character{
		Meta: Meta{
			Sources: []Source{
				{
					Json:         "Malgorgon",
					Abbreviation: "MG",
					Authors:      []string{"Krzysztof"},
					ConvertedBy:  []string{"Krzysztof"},
					Version:      "unknown",
				},
			},
			DateAdded:            now,
			DateLastModified:     now,
			DateLastModifiedHash: fmt.Sprintf("%x", now),
		},
		Monster: []Monster{},
	}

	chars = append(chars, newChar)

	return newChar
}

// parseInt konvertiert einen String zu einem Integer und gibt 0 zurück, wenn die Konvertierung fehlschlägt
func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
