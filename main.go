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
	content  embed.FS
	Monsters []model.Monster
)

func main() {
	filename := ""

	http.HandleFunc("/", handlers.FormHandler(content, &Monsters, filename))
	http.HandleFunc("/submit", handlers.SubmitHandler(content, &chars, &Monsters, filename))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.FS(content))))
	http.HandleFunc("/addMonster", handlers.AddMonster(&Monsters))
	http.HandleFunc("/about", handlers.AboutHandler(content))
	http.HandleFunc("/contact", handlers.ContactHandler(content))
	http.HandleFunc("/monsterTable", handlers.MonsterTableHandler(content, &Monsters))

	log.Print("Server gestartet, erreichbar unter http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
