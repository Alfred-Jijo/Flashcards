package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	flachcards "github.com/epicnotgames/cli-flashcards/src/packages/flashcards"
	utils "github.com/epicnotgames/cli-flashcards/src/packages/utils"
)

func main() {
	//flach cuz naming clash
	flashcards := flachcards.GetFlashcards()
	scanner := bufio.NewScanner(os.Stdin)

	fileName := flag.String("file", "flashcards.json", "Specify the flashcards file name")
	flag.Parse()

	fmt.Println("Welcome to the Flashcard Quiz!")

	utils.ErrHandle(flachcards.LoadFlashcards(*fileName))

	for {
		fmt.Print("Would you like to add a flashcard? (y/n): ")
		scanner.Scan()
		answer := scanner.Text()
		if answer == "y" {
			// Save flashcards to file after adding
			utils.ErrHandle(flachcards.SaveFlashcards(*fileName))
		} else if answer == "n" {
			utils.ClearTerminal()
			break
		} else {
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")
		}
	}

	shuffledFlashcards := flachcards.GetShuffledFlashcards(flashcards)

	flachcards.PrintFlashCards(flashcards, shuffledFlashcards, scanner)

	fmt.Println("Flashcard Quiz completed.")
}
