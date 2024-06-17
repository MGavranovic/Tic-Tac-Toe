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

// var for holding user input
var userInput string

func main() {
	displayField()

	// taking user input
	// TODO: will have
	fmt.Print("Please enter a number 1-9: ")
	fmt.Scan(&userInput)
	fmt.Println(userInput)

	// looping over the PlayingField and assigning user input to the selected field
	for i := 0; i < len(PlayingField); i++ {
		for j := 0; j < len(PlayingField); j++ {
			if PlayingField[i][j] == userInput {
				PlayingField[i][j] = "X"
				displayField()
			}
		}
	}
}

// formatting the playing field and printing it
func displayField() {
	for i := 0; i < len(PlayingField); i++ {
		fmt.Printf("%s\n", strings.Join(PlayingField[i], " "))
	}
}
