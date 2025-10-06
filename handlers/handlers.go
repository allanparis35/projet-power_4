package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"projet-power_4/game"
	"projet-power_4/winner"
)

func ShowBoard(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("public/board.html"))
	t.Execute(w, game.State)
}

func PlayMove(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" || game.State.Winner != "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	colStr := r.FormValue("column")
	col, err := strconv.Atoi(colStr)
	
	if err != nil || col < 0 || col > 6 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if game.PlacePiece(col) {
		// Vérifier s'il y a un gagnant après avoir placé la pièce
		game.State.Winner = winner.CheckWinner(game.State.Board)
		
		// Changer de joueur seulement si personne n'a gagné
		if game.State.Winner == "" {
			game.SwitchPlayer()
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		game.Reset()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}