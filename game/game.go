package game


type GameState struct {
	Board         [6][7]string
	CurrentPlayer string
	Winner        string
}

var State GameState

func Reset() {
	State = GameState{
		Board:         [6][7]string{},
		CurrentPlayer: "red",
		Winner:        "",
	}
}

func PlacePiece(col int) bool {

	for row := 5; row >= 0; row-- {
		if State.Board[row][col] == "" {
			State.Board[row][col] = State.CurrentPlayer
			return true
		}
	}
	return false
}

func SwitchPlayer() {
	if State.CurrentPlayer == "red" {
		State.CurrentPlayer = "yellow"
	} else {
		State.CurrentPlayer = "red"
	}
}