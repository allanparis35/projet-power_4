package winner

func CheckWinner(board [6][7]string) string {
	directions := [][2]int{
		{0, 1},
		{1, 0},
		{1, 1},
		{1,-1},
	}

	for row := 0; row < 6; row++ {
		for col := 0; col < 7; col++ {

			if board[row][col] == "" {
				continue
			}

			player := board[row][col]

			for _, dir := range directions {
				count := 1

				for step := 1; step < 4; step++ {
					r := row + dir[0]*step
					c := col + dir[1]*step

					if r < 0 || r >= 6 || c < 0 || c >= 7 || board[r][c] != player {
						break
					}

					count++
					if count == 4 {
						return player
					}
				}
			}
		}
	}
	return ""
}
