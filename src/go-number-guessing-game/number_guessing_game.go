package main

import (
	"fmt"
	"math/rand/v2"
)

func displayStartUpMessage() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("Select game difficaulty to begin.")
}

func generateRandomNumber() int {
	var randomNumber int
	randomNumber = rand.IntN(100)
	for randomNumber <= 1 {
		randomNumber = rand.IntN(100)
	}
	return randomNumber
}

func startGame() {
	displayStartUpMessage()
	fmt.Println(generateRandomNumber())
}

func main() {
	startGame()
}
