// lage array for items(person, kylling, rev, korn)? bruke conditional statements?
// https://www.golangprograms.com/go-language/arrays.html
// metoder som kan brukes: crossriver() move(item) drop(item)
package event

import "fmt"

// arrays med alle items som kan/må lastes på båten
var items [3]string = [3]string{"høne", "rev", "korn"}

var itemToPut string

// slice test, kun et item
var høne []string = items[:1]

//velge hvilken slice "item" man vil laste på båten
func PutInBoat([]string) string {
	return "[|rev|korn|  vvv\\_høne_pers_/ _________________/øøø]"
}

func PrintArray() {
	fmt.Println(items)
}

func Start() {
	fmt.Printf("[%v|%v|%v|  vvv\\__/ _________________/øøø]", items[0], items[1], items[2])
}

func GetInput() {
	fmt.Printf("\n"+"Skriv hva du vil putte i båten, du kan velge mellom %v, %v og %v"+"\n", items[0], items[1], items[2])

	fmt.Scanln(&itemToPut)

	CheckUserInput()

	if itemToPut == items[0] {
		fmt.Println("\n" + "du har valgt høne")
	}
}

func CheckUserInput() {

	if itemToPut != items[0] && itemToPut != items[1] && itemToPut != items[2] {
		fmt.Printf("Du må velge mellom %v, %v, og %v"+"\n", items[0], items[1], items[2])
	}
}

func FlyttHøneØst() {
	fmt.Println("\n" + "du har valgt høne")
}
