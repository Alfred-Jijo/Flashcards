package types

var flashcards = map[string]string{
	"Question 1": "Answer 1",
	"Question 2": "Answer 2",
	"Question 3": "Answer 3",
	// Add more flashcards as needed
}

func GetFlashcards() map[string]string {
	return flashcards
}
