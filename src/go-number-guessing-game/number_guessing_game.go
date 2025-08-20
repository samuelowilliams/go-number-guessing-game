package main

import "fmt"

func displayStartUpMessage() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("Select game difficaulty to begin.")
}

func startGame() { displayStartUpMessage() }
func main() {
	startGame()
}
