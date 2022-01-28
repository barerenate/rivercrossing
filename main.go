package main

import (
	"fmt"

	"github.com/barerenate/rivercrossing/event"
)

func main() {

	/*event.Start()

	fmt.Println("\n", "hello")*/

	// Println function is used to
	// display output in the next line
	fmt.Println("This is the Rivercrossing game")
	fmt.Println("This is the initial tate")
	event.PrintArray()
	event.Start()
	event.GetInput()

}
