package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	flashcards := map[string]string{
		"Question 1": "Answer 1",
		"Question 2": "Answer 2",
		// Add more flashcards as needed
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Flashcard Quiz!")

	for question, answer := range flashcards {
		fmt.Printf("Q: %s\n", question)
		fmt.Print("Press 'Enter' to reveal the answer...")
		scanner.Scan()
		fmt.Printf("A: %s\n\n", answer)
	}

	fmt.Println("Flashcard Quiz completed. Goodbye!")
}
