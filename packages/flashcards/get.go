package flachcards

import (
	"math/rand"
	"time"

	"github.com/epicnotgames/cli-flashcards/packages/storage"
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
