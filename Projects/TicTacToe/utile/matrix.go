package play

import "fmt"

type TictactoeMatrix struct {
	mat    [3][3]string
	symbol string
}

// Initiate empty Tic Tac Toe Matrix
func (m *TictactoeMatrix) InitiatMatrix() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m.mat[i][j] = "-"
		}
	}

}

// Update the Tic Tac Toe Matrix
func (m *TictactoeMatrix) UpdateMatrix(i int, j int, plyer string) {
	if plyer == "O" {
		m.mat[i][j] = "O"
	} else if plyer == "X" {
		m.mat[i][j] = "X"
	}

}

// Print the Tic Tac Toe Matrix
func (m *TictactoeMatrix) PrintMatrix() {
	for i := 0; i < 3; i++ {
		fmt.Println(m.mat[i])
	}
}

func (m *TictactoeMatrix) CheckPosition(i int, j int) bool {
	var replay bool
	if m.mat[i][j] == "O" || m.mat[i][j] == "X" {
		fmt.Println("This position is taken. Choose another point")
		replay = true
	} else {
		replay = false
	}
	return replay
}

func (m *TictactoeMatrix) CheckAWin(player string, symbol string) bool {
	var awin bool
	if symbol == "O" {
		for i := 0; i <= 2; i++ {
			if m.mat[i][0] == "O" && m.mat[i][1] == "O" && m.mat[i][2] == "O" {
				m.symbol = "O"
			}
		}
		for j := 0; j <= 2; j++ {
			if m.mat[0][j] == "O" && m.mat[1][j] == "O" && m.mat[2][j] == "O" {
				m.symbol = "O"
			}
		}
		if m.mat[0][0] == "O" && m.mat[1][1] == "O" && m.mat[2][2] == "O" {
			m.symbol = "O"
		}
		if m.mat[0][2] == "O" && m.mat[1][1] == "O" && m.mat[2][0] == "O" {
			m.symbol = "O"
		}
	}

	if symbol == "X" {
		for i := 0; i <= 2; i++ {
			if m.mat[i][0] == "X" && m.mat[i][1] == "X" && m.mat[i][2] == "X" {
				m.symbol = "X"
			}
		}
		for j := 0; j <= 2; j++ {
			if m.mat[0][j] == "X" && m.mat[1][j] == "X" && m.mat[2][j] == "X" {
				m.symbol = "X"
			}
		}
		if m.mat[0][0] == "X" && m.mat[1][1] == "X" && m.mat[2][2] == "X" {
			m.symbol = "X"
		}
		if m.mat[0][2] == "X" && m.mat[1][1] == "X" && m.mat[2][0] == "X" {
			m.symbol = "X"
		}
	}

	switch m.symbol {
	case "O":
		fmt.Println("The winer is: ", player)
		awin = true
	case "X":
		fmt.Println("The winer is: ", player)
		awin = true
	default:
		fmt.Println("There are no winer")
		awin = false
	}
	return awin

}
