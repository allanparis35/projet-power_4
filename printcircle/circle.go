package piscine 

 func PrintCircle() {
	var rows = 6
	var columns = 7
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if (i+j)%2 == 0 {
				print("O ")
			} else {
				print(". ")
			}
		}
	}
 }