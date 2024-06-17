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

// TODO: have to add players
var player int = 2

// var for holding user input
var userInput string

// this variable is used for making a do while loop so I can run stuff as long as the program is running or certain conditions
// (win or draw) are met
var running bool = true

func main() {
	for {
		displayField()
		// testing how to access elements of slice
		fmt.Println(PlayingField[0][0])

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
		fmt.Print("Please enter a number 1-9: ")
		fmt.Scan(&userInput)
		fmt.Println(userInput)

		updateField()
		// // looping over the PlayingField and assigning user input to the selected field
		// for i := 0; i < len(PlayingField); i++ {
		// 	for j := 0; j < len(PlayingField); j++ {
		// 		if PlayingField[i][j] == userInput {
		// 			if player == 1 {
		// 				PlayingField[i][j] = "X"
		// 			} else if player == 2 {
		// 				PlayingField[i][j] = "O"
		// 			}
		// 			// displaying the updated field
		// 			// displayField()
		// 		}
		// 	}
		// }
		if running == false {
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
				// displaying the updated field
				// displayField()
			}
		}
	}
}
