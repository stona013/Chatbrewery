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

// main is the entry point of the program.
func main() {
	filename := ""

	// Create a new ServeMux instance
	routes := http.NewServeMux()

	// Register the handlers for different routes
	routes.HandleFunc("/", handlers.FormHandler(content, &Monsters))
	routes.HandleFunc("/submit", handlers.SubmitHandler(content, &chars, &Monsters, filename))
	routes.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.FS(content))))
	routes.HandleFunc("/addMonster", handlers.AddMonster(&Monsters))
	routes.HandleFunc("/main", handlers.MainHandler(content, &Monsters))
	routes.HandleFunc("/about", handlers.AboutHandler(content))
	routes.HandleFunc("/contact", handlers.ContactHandler(content))
	routes.HandleFunc("/monsterTable", handlers.MonsterTableHandler(content, &Monsters))
	// Print the message indicating that 'static' has been included.
	log.Printf("Eingebunden is %v\n", static)

	// Load the CSS file.
	css, err := loadCSS(static)
	if err != nil {
		log.Fatal(err)
	}

	// Add a route for the CSS file
	routes.HandleFunc("/static/darkly_bulmawatch.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		w.Write([]byte(css))
	})

	// Print the message indicating that the server has started.
	log.Print("Server gestartet, erreichbar unter http://localhost:8080")

	// Start the server and listen for incoming requests on port 8080.
	log.Fatal(http.ListenAndServe(":8080", routes))
}

// loadCSS reads the CSS file from the embedded filesystem.
// It takes the content embed.FS as input.
// It returns the content of the CSS file as a string and an error if any.
func loadCSS(content embed.FS) (string, error) {
	// Read the CSS file "static/darkly_bulmawatch.css" from the embedded filesystem
	file, err := content.ReadFile("static/darkly_bulmawatch.css")
	if err != nil {
		return "", err
	}
	// Convert the file content to a string and return
	return string(file), nil
}
