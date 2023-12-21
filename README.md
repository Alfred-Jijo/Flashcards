# Flashcards

## Overview
This program implements a simple flashcard quiz using the command line interface. Users can add flashcards with questions and answers, shuffle them, and quiz themselves. The flashcards are stored in a predefined set.

## Starting Out
- If you want a fresh set of flashcards empty `flashcards.json` to look likt this:
```
{

}
```
At least one flashcards must be made to start the program
```
{
    "Key": { <== Key can be anything as long as it does conflict with other keys 
        "Question":  "...", <== must have a comma
        "Answer":  "..."
    }
}
```
Recommended json:
```
{
    "Flashcard/Question 1": {
        "Question" : "Qusetion...",
        "Answer" : "Answer"
    }
}
```

## Cli Flags

- you can do ```main.exe -file filename.json``` leave for flashcards.json

## How to Use

1. Run the program.
2. Choose whether to add a new flashcard.
3. If adding a flashcard, enter the question and answer.
4. If not adding, proceed to the flashcard quiz.
5. Answer the questions displayed on the screen.
6. The program will clear the terminal after each flashcard.
7. Once the quiz is completed, the program will exit.

## Note
- New sets must be manually made via making a file called filename.json