package main

import (
	"bufio"
	"fmt"
	"os"

	flachcards "github.com/Drill-Byte/cli-flashcards/flashcards"
)

func main() {
	flashcards := flachcards.GetFlashcards()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Flashcard Quiz!")
	fmt.Println("Press 'Enter' to reveal the answer...")
	shuffledFlashcards := flachcards.GetShuffledFlashcards()

	flachcards.PrintFlashCards(flashcards, shuffledFlashcards, scanner)

	fmt.Println("Flashcard Quiz completed. Goodbye!")
}
