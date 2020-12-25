package main

import "fmt"

const mod = 20201227

func findLoop(pubkey int) int {
	subject := 1
	loop := 0

	for subject != pubkey {
		subject *= 7
		subject %= mod
		loop++
	}

	return loop
}

func transform(v, loop int) int {
	subject := 1
	for i := 1; i <= loop; i++ {
		subject *= v
		subject %= mod
	}
	return subject
}

const (
	exCardPublic = 5764801
	exDoorPublic = 17807724

	myCardPublic = 7573546
	myDoorPublic = 17786549
)

func main() {
	card := findLoop(myCardPublic)
	door := findLoop(myDoorPublic)
	key1 := transform(myDoorPublic, card)
	key2 := transform(myCardPublic, door)
	fmt.Println("P1:", key1, key2)
}
