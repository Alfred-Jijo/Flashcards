package main

import (
	"bufio"
	"fmt"
	"os"

	handle "github.com/Drill-Byte/cli-flashcards/errHandling"
	flachcards "github.com/Drill-Byte/cli-flashcards/flashcards"
)

func main() {
	//flach cuz naming clash
	flashcards := flachcards.GetFlashcards()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Flashcard Quiz!")

	for {
		// !TODO: adding only works for curr session
		fmt.Print("Would you like to add a flashcard? (y/n): ")
		scanner.Scan()
		answer := scanner.Text()
		if answer == "y" {
			succes, err := flachcards.AddFlashcard(flashcards, scanner)
			handle.ErrHandle(err)
			handle.SuccessHandle(succes)
		} else if answer == "n" {
			break
		} else {
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")
		}
	}
	shuffledFlashcards := flachcards.GetShuffledFlashcards()

	flachcards.PrintFlashCards(flashcards, shuffledFlashcards, scanner)

	fmt.Println("Flashcard Quiz completed. Goodbye!")
}
