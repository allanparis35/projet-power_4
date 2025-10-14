package main

import (
	"net/http"
	"projet-power_4/game"
	"projet-power_4/handlers"
)

func main() {
	game.Reset()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlers.ShowBoard)
	http.HandleFunc("/play", handlers.PlayMove)
	http.HandleFunc("/reset", handlers.ResetHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}