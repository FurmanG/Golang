package main

import (
	"fmt"
	mem "packages/utile"
)

func main() {
	m1 := new(mem.Member)
	m2 := new(mem.Member)
	fmt.Println("----------------")
	// Set m1 memeber data
	m1.SetaMemeber("jon", "Smith", 33)
	// Get m1 member data
	m1.GetaMemeber()
	fmt.Println("----------------")
	// Set m2 memeber data
	m2.SetaMemeber("Bob", "Cohen", 28)
	// Get m1 member data
	m2.GetaMemeber()

}
