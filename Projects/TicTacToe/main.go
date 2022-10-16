package main

import (
	"fmt"
	m "tictactoe/utile"
)

func main() {
	var awin bool = false
	// Initiate Tic Tac Toe matrix
	currentMatrix := new(m.TictactoeMatrix)
	currentMatrix.InitiatMatrix()
	currentMatrix.PrintMatrix()

	// Initiate players
	player1 := new(m.Player)
	player2 := new(m.Player)
	player1.InitialPlayer("O")
	player2.InitialPlayer("X")

	// Play and update matrix
	counter := 0
	for {
		// First player
		row1, column1, plysymbol1, name1 := player1.Play()
		// Check if a position in the matrix has already been selected
		for currentMatrix.CheckPosition(row1, column1) {
			row1, column1, plysymbol1, name1 = player1.Play()
		}
		// Update current matrix
		currentMatrix.UpdateMatrix(row1, column1, plysymbol1)
		currentMatrix.PrintMatrix()
		// Checking if there is a winner
		awin = currentMatrix.CheckAWin(name1, plysymbol1)
		if awin {
			break
		}
		// Once all positions are marked, the game ends without a winner.
		counter++
		if counter == 9 {
			fmt.Println("Oops!!! no winer")
			break
		}

		// Second player
		row2, column2, plysymbol2, name2 := player2.Play()
		for currentMatrix.CheckPosition(row2, column2) {
			row2, column2, plysymbol2, name2 = player2.Play()
		}
		currentMatrix.CheckPosition(row2, column2)
		currentMatrix.UpdateMatrix(row2, column2, plysymbol2)
		currentMatrix.PrintMatrix()
		awin = currentMatrix.CheckAWin(name2, plysymbol2)
		if awin {
			break
		}
		counter++
		if counter == 9 {
			fmt.Println("Oops!!! no winer")
			break
		}

	}
}
