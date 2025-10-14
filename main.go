package main
import (
    "net/http"
	"projet-power_4/handlers"
	"projet-power_4/table"
)
func main() {

	table.Reset()
	
	http.HandleFunc("/", handlers.ShowBoard)
	http.HandleFunc("/play", handlers.PlayMove)
	http.HandleFunc("/reset", handlers.ResetHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}

