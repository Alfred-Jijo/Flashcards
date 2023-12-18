package flachcards

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/epicnotgames/cli-flashcards/packages/storage"
)

const flashcardsFile = "flashcards.json"

func LoadFlashcards() error {
	file, err := os.Open(flashcardsFile)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&storage.Flashcards); err != nil {
		return err
	}

	fmt.Println("Flashcards loaded successfully.")
	return nil
}
