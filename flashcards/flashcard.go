package flachcards

import (
	"bufio"
	"fmt"
	"math/rand"
	"time"

	"github.com/Drill-Byte/cli-flashcards/storage"
)

func GetFlashcards() map[string]storage.Flashcard {
	return storage.Flashcards
}

func GetShuffledFlashcards(flashcards map[string]storage.Flashcard) []string {
	keys := make([]string, 0, len(flashcards))
	for key := range flashcards {
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
	}
	// for _, question := range shuffledFlashcards {
	// 	fmt.Printf("Q: %s\n", question)
	// 	scanner.Scan()
	// 	fmt.Printf("A: %s\n\n", flashcards[question])
	// 	//fmt.Printf("Q: %s\n\nA: %s\n\n", card.Question, card.Answer)
	// }
}
