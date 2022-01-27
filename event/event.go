// lage array for items(person, kylling, rev, korn)? bruke conditional statements?
// https://www.golangprograms.com/go-language/arrays.html
// metoder som kan brukes: crossriver() move(item) drop(item)
// var item []string = []string{"rev", "høne", "korn", "pers"}
//Structs https://www.youtube.com/watch?v=sPX6ORiyd0o&t=658s
package event

import "fmt"

var items [4]string = [4]string{"høne", "rev", "korn", "pers"}
var høne []string = items[0:1]

func Start() {
	fmt.Printf("[%v|%v|%v|%v  vvv\\__/ _________________/øøø]", items[0], items[1], items[2], items[3])
}

func PutInBoat(høne []string) string {
	return "[|rev|korn|  vvv\\_høne_pers_/ _________________/øøø]"
}

func CheckState() {

}

func CrossRiver() {

}
