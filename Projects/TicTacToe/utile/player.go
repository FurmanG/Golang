package play

import "fmt"

type Player struct {
	name     string
	xoSymbol string
	pmat     [3][3]string
	win      bool
}

// Initial player
func (p *Player) InitialPlayer(xoSymbol string) {
	fmt.Printf("Enter player name: ")
	fmt.Scan(&p.name)
	p.xoSymbol = xoSymbol
}

// In turn, each player marks the matrix with an X or an O.
func (p *Player) Play() (int, int, string, string) {
	var i int
	var j int

	fmt.Printf("%v - Please choose: \nrow: (0 1 2) : ", p.name)

	fmt.Scan(&i)
	fmt.Printf("Column: (0 1 2) : ")

	fmt.Scan(&j)
	return i, j, p.xoSymbol, p.name
}
