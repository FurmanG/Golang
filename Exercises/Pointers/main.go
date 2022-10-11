package main

import "fmt"

func main() {

	// The * operator allows us to access the values held at a location in memory.
	// Think of a book.
	// The & operator would generate a page number,
	// the * operator would tell you what’s on a given page number.

	//Declaring 2 variables
	// They get a space and an address in memory
	i, j := 2, 12

	// By adding an ampersand before the variable we will get the address of the variable
	fmt.Println("The address of i is = ", &i)
	fmt.Println("The address of j is = ", &j) // the & is "The address of"

	p := &i
	// printing p wil the address also
	fmt.Println("The address of i is = ", p)
	// but printing the *p will give the value of i
	// We call it dereferencing
	fmt.Println("The value of i is = ", *p)

	// changing *p will change the value of i
	*p = 4
	fmt.Println("i is now = ", i)
	// the value of j is now stored in *p
	p = &j
	fmt.Println("*p is now = j = ", *p)

	*p = *p / 2
	fmt.Println("The value of j now is = ", *p)

}
