package flachcards

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/epicnotgames/cli-flashcards/packages/storage"
)

const flashcardsFile = "flashcards.json"

func SaveFlashcards() error {
	file, err := os.Create(flashcardsFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(storage.Flashcards); err != nil {
		return err
	}

	fmt.Println("Flashcards saved successfully.")
	return nil
}
