package flachcards

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/epicnotgames/cli-flashcards/packages/storage"
)

// const flashcardsFile = "flashcards.json"

func LoadFlashcards(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&storage.Flashcards); err != nil {
		return err
	}
	fmt.Printf("Loading flashcards from %s\n", fileName)
	fmt.Println("Flashcards loaded successfully.")
	return nil
}
