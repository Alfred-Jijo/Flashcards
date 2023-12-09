package flachcards

import (
	"bufio"
	"fmt"
	"math/rand"
	"time"

	"github.com/Drill-Byte/cli-flashcards/storage"
)

func GetFlashcards() map[string]string {
	return storage.Flashcards
}

func GetShuffledFlashcards(flashcards map[string]string) []string {
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
	flashcards map[string]string,
	shuffledFlashcards []string,
	scanner *bufio.Scanner,
) {
	for _, question := range shuffledFlashcards {
		fmt.Printf("Q: %s\n", question)
		scanner.Scan()
		fmt.Printf("A: %s\n\n", flashcards[question])
	}
}
