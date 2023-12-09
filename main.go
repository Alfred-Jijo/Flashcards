package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Drill-Byte/cli-flashcards/types"
)

func shuffleFlashcards(flashcards map[string]string) []string {
	keys := make([]string, 0, len(flashcards))
	for key := range flashcards {
		keys = append(keys, key)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	return keys
}

func main() {
	flashcards := types.GetFlashcards()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Flashcard Quiz!")

	shuffledFlashcards := shuffleFlashcards(flashcards)

	for _, question := range shuffledFlashcards {
		fmt.Printf("Q: %s\n", question)
		fmt.Print("Press 'Enter' to reveal the answer...")
		scanner.Scan()
		fmt.Printf("A: %s\n\n", flashcards[question])
	}

	fmt.Println("Flashcard Quiz completed. Goodbye!")
}
