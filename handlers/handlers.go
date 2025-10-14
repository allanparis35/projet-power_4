package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"projet-power_4/table"
	"projet-power_4/winner"
)

func ShowBoard(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("public/board.html"))
	t.Execute(w, table.State)
}

func PlayMove(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" || table.State.Winner != "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	colStr := r.FormValue("column")
	col, err := strconv.Atoi(colStr)

	if err != nil || col < 0 || col > 6 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

	if table.PlacePiece(col) {
		table.State.Winner = winner.CheckWinner(table.State.Board)

		if table.State.Winner == "" {
			table.SwitchPlayer()
		}
	}
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		table.Reset()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
