package handlers

import (
	"html/template"
	"net/http"

	flachcards "github.com/epicnotgames/cli-flashcards/src/packages/flashcards"
	"github.com/epicnotgames/cli-flashcards/src/packages/storage"
	"github.com/epicnotgames/cli-flashcards/src/packages/utils"
)

// func IndexHandler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "web/templates/index.html")
// }

func FlashcardsHandler(w http.ResponseWriter, r *http.Request) {
	flashcards := flachcards.GetFlashcards()
	shuffledFlashcards := flachcards.GetShuffledFlashcards(flashcards)

	tmpl, err := template.ParseFiles("templates/index.html")
	utils.ErrHandle(err)

	data := struct {
		Flashcards map[string]storage.Flashcard
		Shuffled   []string
	}{
		Flashcards: flashcards,
		Shuffled:   shuffledFlashcards,
	}

	err = tmpl.Execute(w, data)
	utils.ErrHandle(err)
}
