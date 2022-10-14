package play

import "fmt"

type TictactoeMatrix struct {
	mat [3][3]string
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
