package main

import (
	mem "tictactoe/utile"
)

func main() {
	var awin bool = false
	// Initiate Tic Tac Toe matrix
	currentMatrix := new(mem.TictactoeMatrix)
	currentMatrix.InitiatMatrix()
	currentMatrix.PrintMatrix()

	// Initiate players
	p1 := new(mem.Player)
	p2 := new(mem.Player)
	p1.InitialPlayer("Joe", "O")
	p2.InitialPlayer("Bob", "X")

	// Play and update matrix
	for {
		i1, j1, plysymbol1, name1 := p1.Play()
		for currentMatrix.CheckPosition(i1, j1) {
			i1, j1, plysymbol1, name1 = p1.Play()
		}
		currentMatrix.UpdateMatrix(i1, j1, plysymbol1)
		currentMatrix.PrintMatrix()
		awin = currentMatrix.CheckAWin(name1, plysymbol1)
		if awin {
			break
		}

		i2, j2, plysymbol2, name2 := p2.Play()
		for currentMatrix.CheckPosition(i2, j2) {
			i2, j2, plysymbol2, name2 = p2.Play()
		}
		currentMatrix.CheckPosition(i2, j2)
		currentMatrix.UpdateMatrix(i2, j2, plysymbol2)
		currentMatrix.PrintMatrix()
		awin = currentMatrix.CheckAWin(name2, plysymbol2)
		if awin {
			break
		}

	}
}
