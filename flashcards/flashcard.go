package types

import (
	"math/rand"
	"time"
)

var flashcards = map[string]string{
	"Question 1": "Answer 1",
	"Question 2": "Answer 2",
	"Question 3": "Answer 3",
	// Add more flashcards as needed
	"Question 4": "Answer 4",
	"Question 5": "Answer 5",
	"Question 6": "Answer 6",
}

func ShuffleFlashcards(flashcards map[string]string) []string {
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

func GetFlashcards() map[string]string {
	return flashcards
}

func GetShuffledFlashcards() []string {
	return ShuffleFlashcards(flashcards)
}
