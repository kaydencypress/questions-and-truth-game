package main

import (
	"example.com/questions-and-truth/cards"
	"example.com/questions-and-truth/questions"
)

func main() {
	// Choose random cards for hand from standard deck of 52 cards and initialize questions
	deck := cards.NewDeck()
	deck.ShuffleDeck()
	questionSet := questions.InitializeQuestions(deck)
	hand, err := cards.NewHand(&deck, 8)

	if err != nil {
		panic(err)
	}

	// Select order of cards in hand
	hand.ReorderHand()

	// TODO: menu to select a question or guess truth & loop

	// If Question:
	// Prompt user to select a question and print answer
	err = questions.SelectQuestionAndGetAnswer(questionSet, hand)
	if err != nil {
		panic(err)
	}

	// If truth:
	// TODO: Prompt user for guess
	// TODO: Check if guess is correct
}
