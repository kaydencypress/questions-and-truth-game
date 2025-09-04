package main

import (
	"fmt"

	"example.com/questions-and-truth/cards"
	"example.com/questions-and-truth/menus"
	"example.com/questions-and-truth/questions"
	"example.com/questions-and-truth/truth"
)

func main() {
	// Choose random cards for hand from standard deck of 52 cards and initialize questions
	cards.GameDeck.ShuffleDeck()
	questionSet := questions.InitializeQuestions()
	hand, err := cards.NewHand(&cards.GameDeck, 8)

	if err != nil {
		panic(err)
	}

	// Select order of cards in hand
	hand.ReorderHand()

	// menu to select a question or guess truth
	for {
		selection, err := menus.SelectQuestionOrTruth()

		if err != nil {
			fmt.Println(err)
			continue
		}

		// Exit program
		if selection == 3 {
			return
		}

		// "Question" selected
		if selection == 1 {
			// Prompt user to select a question and print answer
			err = questions.SelectQuestionAndGetAnswer(questionSet, hand)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		// "Truth" selected
		if selection == 2 {
			// Prompt user for guess
			guess, err := truth.GetUserGuess(hand)
			if err != nil {
				fmt.Println(err)
				continue
			}
			// Check if guess is correct
			isGuessCorrect := truth.IsGuessCorrect(guess, hand)

			if isGuessCorrect {
				fmt.Println("Correct - You won!")
				break
			} else {
				fmt.Println("Incorrect - Try again.")
			}
		}
	}

}
