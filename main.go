package main

import (
	"fmt"

	"github.com/barerenate/rivercrossing/metoder"
)

func main() {

	fmt.Println("Rivercrossing" + "\n" + "Status ved start:")
	metoder.InitState()
	metoder.MoveHØ()
	metoder.MoveRØ()

}
