package storage

type Flashcard struct {
	Question string
	Answer   string
}

var Flashcards = map[string]Flashcard{
	"Question_One": {
		Question: "What is a Quantity",
		Answer:   "In S.I. a quantity is represented by a number and a unit, (e.g. m = 3.0 kg).",
	},
	"Question_Two": {
		Question: "What is a Vector",
		Answer:   "A vector is a quantity that has magnitude and direction.",
	},
	"Question_Three": {
		Question: "Newton's First Law",
		Answer:   "A body will remain in uniform motion in a straight line unless acted on by an external force.",
	},
	"Question_Four": {
		Question: "Newton's Second Law",
		Answer:   "The rate of change of a body's momentum is directly proportional to the force active upon it.",
	},
	// Add more flashcards as needed...
}
