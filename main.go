package main

import (
	"ddServer/handlers"
	"ddServer/model"
	"embed"
	"fmt"
	"net/http"
	"sync"
)

var (
	mu    sync.Mutex
	chars []model.Character
	//go:embed templates/*.html
	//go:embed images/*
	content  embed.FS
	Monsters []model.Monster
)

func main() {
	filename := ""

	http.HandleFunc("/", handlers.FormHandler(content, filename))
	http.HandleFunc("/submit", handlers.SubmitHandler(content, &chars, &Monsters, filename))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.FS(content))))
	http.HandleFunc("/addMonster", handlers.AddMonster(&Monsters))
	http.HandleFunc("/about", handlers.AboutHandler(content))
	http.HandleFunc("/contact", handlers.ContactHandler(content))

	fmt.Println("Server gestartet, erreichbar unter http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
