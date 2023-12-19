package flachcards

import (
	"os"

	"github.com/Drill-Byte/cli-flashcards/packages/utils"
)

//a function that clears everything in the flashcards.json

func InitFlashcards() error {
	file, err := os.Create(flashcardsFile)
	utils.ErrHandle(err)
	defer file.Close()

	return nil
}
