package main

import (
	"bufio"
	"fmt"
	"os"

	flachcards "github.com/Drill-Byte/cli-flashcards/flashcards"
)

func main() {
	//flach cuz naming clash
	flashcards := flachcards.GetFlashcards()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Flashcard Quiz!")
	for {
		shuffledFlashcards := flachcards.GetShuffledFlashcards()

		flachcards.PrintFlashCards(flashcards, shuffledFlashcards, scanner)

		var input = scanner.Text()
		if input == "quit" {
			break
		}
	}

	fmt.Println("Flashcard Quiz completed. Goodbye!")
}
