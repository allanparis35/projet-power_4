package handlers

import (
	"html/template"
	"net/http"
	"projet-power_4/table"
)

func ShowBoard(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("public/board.html"))
	t.Execute(w, table.State)
}
func ResetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		table.Reset()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}