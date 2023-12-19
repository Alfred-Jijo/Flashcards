package main

import (
	"bufio"
	"fmt"
	"os"

	flachcards "github.com/Drill-Byte/cli-flashcards/packages/flashcards"
	utils "github.com/Drill-Byte/cli-flashcards/packages/utils"
)

func main() {
	//flach cuz naming clash
	flashcards := flachcards.GetFlashcards()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Flashcard Quiz!")

	utils.ErrHandle(flachcards.LoadFlashcards())

	for {
		fmt.Print("Would you like to add a flashcard? (y/n): ")
		scanner.Scan()
		answer := scanner.Text()
		if answer == "y" {
			// Save flashcards to file after adding
			utils.ErrHandle(flachcards.SaveFlashcards())
		} else if answer == "n" {
			utils.ClearTerminal()
			break
		} else if answer == "init" {
			utils.ErrHandle(flachcards.InitFlashcards())
		} else {
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")
		}
	}

	shuffledFlashcards := flachcards.GetShuffledFlashcards(flashcards)

	flachcards.PrintFlashCards(flashcards, shuffledFlashcards, scanner)

	fmt.Println("Flashcard Quiz completed. Goodbye!")
}
