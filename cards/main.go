package main

import "fmt"

func main() {

	card := pickACard()
	fmt.Println(card)
}

func pickACard() string {
	return "King"
}
