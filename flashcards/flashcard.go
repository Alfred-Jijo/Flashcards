package flachcards

import (
	"bufio"
	"fmt"
	"math/rand"
	"time"

	"github.com/Drill-Byte/cli-flashcards/storage"
	"github.com/Drill-Byte/cli-flashcards/utils"
)

func GetFlashcards() map[string]storage.Flashcard {
	return storage.Flashcards
}

func GetShuffledFlashcards(flashcards map[string]storage.Flashcard) []string {
	keys := make([]string, 0, len(flashcards))
	for key := range storage.Flashcards {
		keys = append(keys, key)
	}

	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	return keys
}

func PrintFlashCards(
	flashcards map[string]storage.Flashcard,
	shuffledFlashcards []string,
	scanner *bufio.Scanner,
) {
	for key, card := range flashcards {
		fmt.Printf("%s\n", key)
		fmt.Printf("Q: %s\n", card.Question)
		scanner.Scan()
		fmt.Printf("A: %s\n\n", card.Answer)
		scanner.Scan()
		utils.ClearTerminal()
	}
}

func AddFlashcard(flashcards map[string]storage.Flashcard, scanner *bufio.Scanner) error {

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
