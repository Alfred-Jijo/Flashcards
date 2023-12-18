package flachcards

import (
	"bufio"
	"fmt"

	"github.com/Drill-Byte/cli-flashcards/packages/storage"
)

func AddFlashcard(
	flashcards map[string]storage.Flashcard,
	scanner *bufio.Scanner,
) error {

	var key = fmt.Sprintf("Question_%d", len(flashcards)+1)

	if _, exists := flashcards[key]; exists {
		return fmt.Errorf("flashcard with the key '%s' already exists", key)
	}

	fmt.Print("Enter the question for the new flashcard: ")
	scanner.Scan()
	question := scanner.Text()

	fmt.Print("Enter the answer for the new flashcard: ")
	scanner.Scan()
	answer := scanner.Text()

	newFlashcard := storage.Flashcard{
		Question: question,
		Answer:   answer,
	}

	flashcards[key] = newFlashcard

	fmt.Println("Flashcard added successfully!")
	return nil
}
