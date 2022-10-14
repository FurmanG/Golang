package main

import (
	mem "tictactoe/utile"
)

func main() {
	// Initiate Tic Tac Toe matrix
	f := new(mem.TictactoeMatrix)
	f.InitiatMatrix()
	f.PrintMatrix()

	// Initiate players
	p1 := new(mem.Player)
	p2 := new(mem.Player)
	p1.InitialPlayer("Joe", "O")
	p2.InitialPlayer("Bob", "X")

	// Play and update matrix
	i1, j1, plyer1 := p1.Play()
	f.UpdateMatrix(i1, j1, plyer1)
	f.PrintMatrix()

	i2, j2, plyer2 := p2.Play()
	f.UpdateMatrix(i2, j2, plyer2)
	f.PrintMatrix()

}
