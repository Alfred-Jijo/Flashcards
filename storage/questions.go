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
	"Question_Five": {
		Question: "Newton's Third Law",
		Answer:   "If a body A exerts a force on a body B, then B exerts an equal and opposite force on A that is of the same type.",
	},
	"Question_Six": {
		Question: "Resolving a Vector into components in paticular directions",
		Answer: `Resolving a  vector into components in particular directions
		This means finding vectors (the so-called components) in these directions,
		which add together vectorially to make the original vector, and so,
		together, are equivalent to this vector.`,
	},
	"Question_Seven": {
		Question: "Density of a material",
		Answer: `Density is Mass / Volume 
		The Mass per Unit volume is kgm^-3
		in which mass and volume apple to any sample of the material`,
	},
	// Add more flashcards as needed...
}
