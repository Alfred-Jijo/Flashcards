## Package Structure

### `flashcards`

#### `AddFlashcard`
- Adds a new flashcard to the collection.
- Checks for duplicate flashcards based on a generated key.

#### `GetFlashcards`
- Retrieves the predefined set of flashcards.

#### `GetShuffledFlashcards`
- Shuffles the keys of the flashcards for a randomized quiz experience.

#### `PrintFlashCards`
- Prints flashcards in a randomized order.
- Clears the terminal for a better user experience.

### `storage`

#### `Flashcard`
- Defines the structure of a flashcard with a question and an answer.

#### `Flashcards`
- Predefined set of flashcards.

### `utils`

#### `runCmd`
- Executes shell commands.

#### `ClearTerminal`
- Clears the terminal screen based on the operating system.

#### `ErrHandle`
- Handles and prints errors.
