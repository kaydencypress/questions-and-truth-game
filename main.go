package main

import (
	"fmt"

	"example.com/questions-and-truth/cards"
	"example.com/questions-and-truth/questions"
)

func main() {
	// Choose random cards for hand from standard deck of 52 cards
	deck := cards.NewDeck()
	deck.ShuffleDeck()
	hand, err := cards.NewHand(&deck, 8)

	if err != nil {
		panic(err)
	}

	// Select order of cards in hand
	hand.ReorderHand()

	questionSet := InitializeQuestions(deck)
	// TODO: menu to select a question or guess truth
	selectedQuestion := questionSet[0]

	// If Question:
	// Prompt for values to substitute into question
	selectedQuestion.DisplayAndPromptIfNeeded()
	// TODO: Determine answer

	// If truth:
	// TODO: Prompt user for guess
	// TODO: Check if guess is correct
}

func InitializeQuestions(deck cards.Deck) []questions.Question {
	suitPromptText := ""
	for _, suit := range deck.Suits {
		suitPromptText += fmt.Sprintf("%d. %s\n", suit.ID, suit.Name)
	}
	suitPromptText += "Select suit (1-4): "
	suitPrompt := []questions.Prompt{{Text: suitPromptText, MaxValue: 4}}
	q1 := questions.NewQuestion("Position", "Which positions contain cards of the specified suit?", suitPrompt)
	return []questions.Question{
		q1,
	}
}
