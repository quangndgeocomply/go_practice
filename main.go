package main

import "fmt"

func checkPointer(arr [5]int) [5]int {
	arr[0] = 1
	return arr
}

func main() {
	// Struct
	quang := person{
		firstName: "Quang",
		lastName:  "Nguyen",
		contactInfo: contactInfo{
			email:   "quang@gmail.com",
			zipCode: 94000,
		},
	}
	pointer := &quang
	pointer.updateName("quangnd")
	pointer.print()

	// shortcut: quang.updateName("quangnd")

	// Map
	colors := createMap()

	colors["yellow"] = "laksjdf"
	delete(colors, "yellow")

	printMap(colors)

	// Interface
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	var balance = [5]int{1000, 2, 3, 17, 50}
	checkPointer(balance)
	fmt.Println(balance)
}
