package main

import (
	"ddServer/handlers"
	"ddServer/model"
	"embed"
	"log"
	"net/http"
	"sync"
)

var (
	mu    sync.Mutex
	chars []model.Character
	//go:embed templates/*.html
	//go:embed images/*
	content embed.FS
	//go:embed static/*
	static   embed.FS
	Monsters []model.Monster
)

func main() {
	filename := ""
	log.Printf("Eingebunden is %v\n", static)

	http.HandleFunc("/", handlers.FormHandler(content, &Monsters, filename))
	http.HandleFunc("/submit", handlers.SubmitHandler(content, &chars, &Monsters, filename))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.FS(content))))
	http.HandleFunc("/addMonster", handlers.AddMonster(&Monsters))
	http.HandleFunc("/main", handlers.MainHandler(content, &Monsters))
	http.HandleFunc("/about", handlers.AboutHandler(content))
	http.HandleFunc("/contact", handlers.ContactHandler(content))
	http.HandleFunc("/monsterTable", handlers.MonsterTableHandler(content, &Monsters))

	// Lade die CSS-Datei
	css, err := loadCSS(static)
	if err != nil {
		log.Fatal(err)
	}

	// Füge eine Route für die CSS-Datei hinzu
	http.HandleFunc("/static/darkly_bulmawatch.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		w.Write([]byte(css))
	})

	log.Print("Server gestartet, erreichbar unter http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// loadCSS liest die CSS-Datei aus dem eingebetteten Dateisystem.
func loadCSS(content embed.FS) (string, error) {
	file, err := content.ReadFile("static/darkly_bulmawatch.css")
	if err != nil {
		return "", err
	}
	return string(file), nil
}
