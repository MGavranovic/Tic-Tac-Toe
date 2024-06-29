package main

import (
	"fmt"
	"strings"
)

// playing field
var PlayingField = [][]string{
	{"| ", " ", " | ", " ", " | ", " ", " |"},
	{"| ", "1", " | ", "2", " | ", "3", " |"},
	{"| ", "-", " | ", "-", " | ", "-", " |"},
	{"| ", "4", " | ", "5", " | ", "6", " |"},
	{"| ", "-", " | ", "-", " | ", "-", " |"},
	{"| ", "7", " | ", "8", " | ", "9", " |"},
	{"| ", " ", " | ", " ", " | ", " ", " |"},
}

// initial field used for resetting
var initialField = [][]string{
	{"| ", " ", " | ", " ", " | ", " ", " |"},
	{"| ", "1", " | ", "2", " | ", "3", " |"},
	{"| ", "-", " | ", "-", " | ", "-", " |"},
	{"| ", "4", " | ", "5", " | ", "6", " |"},
	{"| ", "-", " | ", "-", " | ", "-", " |"},
	{"| ", "7", " | ", "8", " | ", "9", " |"},
	{"| ", " ", " | ", " ", " | ", " ", " |"},
}

// used for resetting the field
var gameDone bool = false

// TODO: have to add players
var player int = 2

// var for holding user input
var userInput string

// this variable is used for making a do while loop so I can run stuff as long as the program is running or certain conditions
// (win or draw) are met
var running bool = true

// this vars will be used for checking inputs
var inputValid bool
var possibleInputSlice = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var initialPossibleInputSlice = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

// this var will be used for draw condition
var turns int = 0

// this var used for holding the score value
type Score struct {
	X, O int
}

// declaring score
var score = Score{0, 0}

func main() {
	for running {
		if gameDone {
			clearScreen()
			displayField()
			fmt.Printf("X:%v - O:%v\n", score.X, score.O)
			fmt.Printf("Hit 'r' if you just want to reset the playing field\nHit 'R' to reset both the field and the score:")
			fmt.Scan(&userInput)
			if userInput == "r" {
				resetField()
				gameDone = false
			} else if userInput == "R" {
				resetScore()
				resetField()
				gameDone = false
			}
			continue
		}

		clearScreen()
		displayField()
		checkVictoryCondition()

		if gameDone {
			continue
		}

		// players logic
		if player == 1 {
			player = 2
		} else if player == 2 {
			player = 1
		}

		// taking user input
		// TODO: will have implement error handling for user input
		// TODO: selectable options can be stored in an array
		// if an item doesn't exist in that array, it can't be selected
		for {
			fmt.Print("Please enter a number 1-9: ")
			fmt.Scan(&userInput)
			// fmt.Println(userInput)

			inputValid = false
			for i := 0; i < len(possibleInputSlice); i++ {
				if possibleInputSlice[i] == userInput {
					inputValid = true
					break
				}
			}
			if inputValid {
				break
			} else {
				fmt.Println("Invalid input! Please select a number in range 1-9 and press ENTER")
			}
		}

		updateField()
		if !running {
			break
		}
	}
}

// formatting the playing field and printing it
func displayField() {
	for i := 0; i < len(PlayingField); i++ {
		fmt.Printf("%s\n", strings.Join(PlayingField[i], " "))
	}
}

func updateField() {
	// looping over the PlayingField and assigning user input to the selected field
	for i := 0; i < len(PlayingField); i++ {
		for j := 0; j < len(PlayingField); j++ {
			if PlayingField[i][j] == userInput {
				if player == 1 {
					PlayingField[i][j] = "X"
				} else if player == 2 {
					PlayingField[i][j] = "O"
				}

				// remove the element from the slice
				possibleInputSlice = updatePossibleEntries(possibleInputSlice, userInput)
				turns++
			}
		}
	}
}

// clears the console so the field stays on top and the input field gets reset
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// used for removing removing elements so we can have error handling with user input
func updatePossibleEntries(slice []string, el string) []string {
	for i := 0; i < len(slice); i++ {
		if slice[i] == el {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	// returns a new slice with an element removed
	return slice
}

// function for checking victory conditions and displaying the victory/draw message
func checkVictoryCondition() {
	playerSign := []string{"X", "O"}
	for i := 0; i < len(playerSign); i++ {
		if PlayingField[1][1] == playerSign[i] && PlayingField[1][3] == playerSign[i] && PlayingField[1][5] == playerSign[i] ||
			PlayingField[3][1] == playerSign[i] && PlayingField[3][3] == playerSign[i] && PlayingField[3][5] == playerSign[i] ||
			PlayingField[5][1] == playerSign[i] && PlayingField[5][3] == playerSign[i] && PlayingField[5][5] == playerSign[i] ||
			PlayingField[1][1] == playerSign[i] && PlayingField[3][1] == playerSign[i] && PlayingField[5][1] == playerSign[i] ||
			PlayingField[1][3] == playerSign[i] && PlayingField[3][3] == playerSign[i] && PlayingField[5][3] == playerSign[i] ||
			PlayingField[1][5] == playerSign[i] && PlayingField[3][5] == playerSign[i] && PlayingField[5][5] == playerSign[i] ||
			PlayingField[1][1] == playerSign[i] && PlayingField[3][3] == playerSign[i] && PlayingField[5][5] == playerSign[i] ||
			PlayingField[1][5] == playerSign[i] && PlayingField[3][3] == playerSign[i] && PlayingField[5][1] == playerSign[i] {

			if playerSign[i] == "X" {
				fmt.Println("Player 1 won")
				gameDone = true
				// update the score
				score.X++
			} else {
				fmt.Println("Player 2 won")
				gameDone = true
				// update the score
				score.O++
			}

		} else if turns == 10 {
			fmt.Println("DRAW!")
			gameDone = true
			return
		}
	}
}

// func for resetting the field after the game is done
func resetField() {
	// have to make it as a copy as if use as regular assignment, it will point to the same memory location
	// changing one will affect both
	// PlayingField = initialField
	// possibleInputSlice = initialPossibleInputSlice
	PlayingField = make([][]string, len(initialField))
	for i := range initialField {
		PlayingField[i] = make([]string, len(initialField[i]))
		copy(PlayingField[i], initialField[i])
	}
	possibleInputSlice = make([]string, len(initialPossibleInputSlice))
	copy(possibleInputSlice, initialPossibleInputSlice)
	turns = 0
	player = 2
}

// very simple func for resetting the score
// I'm making it as a function as it makes the code more readable
func resetScore() {
	score = Score{0, 0}
}
