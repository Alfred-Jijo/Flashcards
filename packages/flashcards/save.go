package flachcards

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/epicnotgames/cli-flashcards/packages/storage"
)

const flashcardsFile = "flashcards.json"

func SaveFlashcards(fileName string) error {
	file, err := os.Create(fileName)
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
