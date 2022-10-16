package play

import "fmt"

type Player struct {
	name     string
	xoSymbol string
	pmat     [3][3]string
	win      bool
}

// Initial player
func (p *Player) InitialPlayer(name string, xoSymbol string) {
	p.name = name
	p.xoSymbol = xoSymbol
}

// A player in turn marks the matrix with X or O.
func (p *Player) Play() (int, int, string, string) {
	var i int
	var j int

	fmt.Printf("%v ('%v') Please choose row: (0 1 2) : ", p.name, p.xoSymbol)

	fmt.Scan(&i)
	fmt.Printf("%v ('%v') Please choose column: (0 1 2) : ", p.name, p.xoSymbol)

	fmt.Scan(&j)
	fmt.Printf("You choose row: %v and column: %v\n", i, j)

	return i, j, p.xoSymbol, p.name
}
