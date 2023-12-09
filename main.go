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
	shuffledFlashcards := flachcards.GetShuffledFlashcards(flashcards)

	flachcards.PrintFlashCards(flashcards, shuffledFlashcards, scanner)

	fmt.Println("Flashcard Quiz completed. Goodbye!")
}
