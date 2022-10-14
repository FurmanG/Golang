package play

import "fmt"

type Player struct {
	name     string
	xoSymbol string
	pmat     [3][3]string
}

// Initial player
func (p *Player) InitialPlayer(name string, xoSymbol string) {
	p.name = name
	p.xoSymbol = xoSymbol
}

// A player in turn marks the matrix with X or O.
func (p *Player) Play() (int, int, string) {
	fmt.Printf("%v Please choose row: 0 1 or 2 : ", p.name)
	var i int
	fmt.Scan(&i)
	fmt.Printf("%v Please choose row: 0 1 or 2 : ", p.name)
	var j int
	fmt.Scan(&j)
	fmt.Printf("You choose row: %v and column: %v\n", i, j)
	return i, j, p.xoSymbol
}
