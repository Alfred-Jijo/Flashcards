package flachcards

import (
	"bufio"
	"fmt"
	"math/rand"
	"time"
)

var flashcards = map[string]string{
	"What is a Quantity":  "In S.I. a quantity is represented by a number and a unit, (e.g. m = 3.0 kg).",
	"What is a Vector":    "A vector is a quantity that has magnitude and direction.",
	"Newton's First Law":  "A body will remain in uniform motion in a straight line unless acted on by an external force.",
	"Newton's Second Law": "The rate of change of a body's momentum is directly proportional to the force active upon it.",
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

func PrintFlashCards(
	flashcards map[string]string,
	shuffledFlashcards []string,
	scanner *bufio.Scanner,
) {
	for _, question := range shuffledFlashcards {
		fmt.Printf("Q: %s\n", question)
		fmt.Print("Press 'Enter' to reveal the answer...")
		scanner.Scan()
		fmt.Printf("A: %s\n\n", flashcards[question])
	}
}
