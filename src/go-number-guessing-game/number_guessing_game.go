package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"time"
)

type Highscore struct {
	Easy   int
	Medium int
	Hard   int
}

func checkFile() {
	filename := "highscore.json"
	_, err := os.Stat(filename)
	if err != nil {
		_, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		highscores := Highscore{}
		dataBytes, err := json.Marshal(highscores)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(filename, dataBytes, 0666)
		if err != nil {
			panic(err)
		}
	}
}

func updateHighscore(highscore Highscore) {
	filename := "highscore.json"
	dataBytes, err := json.Marshal(highscore)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(filename, dataBytes, 0666)
	if err != nil {
		panic(err)
	}
}

func loadHighscores() Highscore {
	filename := "highscore.json"
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var highscores Highscore

	err = json.Unmarshal(file, &highscores)
	if err != nil {
		panic(err)
	}
	return highscores
}

func displayStartUpMessage() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("Select game difficulty to begin.")
}

func generateRandomNumber() int {
	var randomNumber int
	randomNumber = rand.IntN(100)
	for randomNumber <= 1 {
		randomNumber = rand.IntN(100)
	}
	return randomNumber
}

func selectDifficulty() string {
	var choice int
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
	fmt.Println()
	fmt.Print("Enter your choice: ")
	fmt.Scanf("%d", &choice)
	fmt.Println("User's choice: ", choice)

	for choice != 1 && choice != 2 && choice != 3 {
		fmt.Println("You selected a difficulty that does not exist. Please try again.")
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d", &choice)
	}

	if choice == 3 {
		return "Hard"
	}
	if choice == 2 {
		return "Medium"
	}
	return "Easy"
}

func gameLogic(difficulty string, highscores Highscore) {
	var chances int
	var guessedNumber int
	attempts := 0
	randomNumber := generateRandomNumber()
	fmt.Println("The randomNumber", randomNumber)
	var t1 time.Time
	var t2 time.Time
	var timeFormat time.Time
	fmt.Printf("Great! You have selected the %s difficulty level.\n", difficulty)
	fmt.Println("Let's start the game!")
	fmt.Println()

	switch difficulty {
	case "Easy":
		chances = 10
	case "Medium":
		chances = 5
	case "Hard":
		chances = 3
	}
	t1 = time.Now()
	for chances > 0 {
		fmt.Print("Enter your guess: ")
		fmt.Scanf("%d", &guessedNumber)
		attempts++
		chances--

		if guessedNumber == randomNumber {
			t2 = time.Now()
			timeTaken := t2.Sub(t1)
			timeFormat = time.Time{}.Add(timeTaken)

			switch difficulty {
			case "Easy":
				if highscores.Easy == 0 {
					highscores.Easy = attempts
					updateHighscore(highscores)
				} else {
					if highscores.Easy > attempts {
						fmt.Printf("Ding Ding Ding : New Highscore on Easy Ddifficulty: %d attempts\n", attempts)
						highscores.Easy = attempts
						updateHighscore(highscores)
					}
				}

			case "Medium":
				if highscores.Medium == 0 {
					highscores.Medium = attempts
					updateHighscore(highscores)
				} else {
					if highscores.Medium > attempts {
						fmt.Printf("Ding Ding Ding : New Highscore on Medium Ddifficulty: %d attempts\n", attempts)
						highscores.Medium = attempts
						updateHighscore(highscores)
					}
				}
			case "Hard":
				if highscores.Hard == 0 {
					updateHighscore(highscores)
					highscores.Hard = attempts
				} else {
					if highscores.Hard > attempts {
						fmt.Printf("Ding Ding Ding : New Highscore on Hard Ddifficulty: %d attempts\n", attempts)
						updateHighscore(highscores)
						highscores.Hard = attempts
					}
				}
			}

			fmt.Printf("Congratulations! You guessed the correct number in %d attempts\n", attempts)
			fmt.Println("Time Taken: ", timeFormat.Format("15:04:05"))
			return
		}

		if randomNumber > guessedNumber {
			fmt.Printf("Incorrect! The number is greater than %d.\n", guessedNumber)
		}

		if randomNumber < guessedNumber {
			fmt.Printf("Incorrect! The number is lesser than %d.\n", guessedNumber)
		}

	}
	t2 = time.Now()
	timeTaken := t2.Sub(t1)
	timeFormat = time.Time{}.Add(timeTaken)

	fmt.Printf("The correct number was %d\n", randomNumber)
	fmt.Println("Time Taken: ", timeFormat.Format("15:04:05"))
}

func startGame() {
	playing := true
	var playAgain int
	checkFile()
	highscores := loadHighscores()

	fmt.Println("Loaded Highscores")
	fmt.Println(highscores)
	displayStartUpMessage()
	for playing {
		difficulty := selectDifficulty()
		gameLogic(difficulty, highscores)
		fmt.Println("New Game: Try Again.")
		fmt.Println("1. Yes")
		fmt.Println("2. No")
		fmt.Print("Enter your choice(1,2): ")
		fmt.Scanf("%d", &playAgain)
		if playAgain == 1 {
			playing = true
			fmt.Println("New Game...")
		}

		if playAgain == 2 {
			playing = false
		}

	}
}

func main() {
	startGame()
}
