package metoder

import (
	"fmt"
)

// slices som kan holde alle mine elemeter
var itemsV []string = []string{"rev", "korn", "høne"}
var båt []string
var itemsØ []string

//printer ut status for hvordan hvor items befinner seg
func PrintState(itemsV []string, båt []string, itemsØ []string) {
	fmt.Println(itemsV, båt, itemsØ)
}

func InitState() {
	PrintState(itemsV, båt, itemsØ)
}

//flytt høne til østsiden av elven
func MoveHØ() {
	itemsØ = append(itemsØ, "høne")
	itemsV = append(itemsV[:2], itemsV[3:]...)
	PrintState(itemsV, båt, itemsØ)
}

// flytt høne til vestsiden av elven
func MoveHV() {
	itemsØ = append(itemsØ[:0], itemsØ[1:]...)
	itemsV = append(itemsV, "høne")
	PrintState(itemsV, båt, itemsØ)
}

// flytt reven til østsiden
func MoveRØ() {
	if contains(itemsØ, "høne") {
		fmt.Println("Error")
	} else {
		itemsØ = append(itemsØ, "rev")
		itemsV = append(itemsV[:0], itemsV[1:]...)
		PrintState(itemsV, båt, itemsØ)
	}

}

// funksjon som kan sjekke om en slice inneholder en streng
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func PutH() {

}

func PutR() {

}

func PutK() {

}
