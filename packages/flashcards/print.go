package flachcards

import (
	"bufio"
	"fmt"

	"github.com/Drill-Byte/cli-flashcards/packages/storage"
	"github.com/Drill-Byte/cli-flashcards/packages/utils"
)

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
