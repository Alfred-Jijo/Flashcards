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
	shuffledFlashcards := flachcards.GetShuffledFlashcards()

	for _, question := range shuffledFlashcards {
		fmt.Printf("Q: %s\n", question)
		fmt.Print("Press 'Enter' to reveal the answer...")
		scanner.Scan()
		fmt.Printf("A: %s\n\n", flashcards[question])
	}

	fmt.Println("Flashcard Quiz completed. Goodbye!")
}
