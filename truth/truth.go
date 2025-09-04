package truth

import (
	"fmt"

	"example.com/questions-and-truth/cards"
	"example.com/questions-and-truth/menus"
)

func IsGuessCorrect(guess, hand cards.Hand) bool {
	for i, guessedCard := range guess {
		if guessedCard != hand[i] {
			return false
		}
	}
	return true
}

func GetUserGuess(hand cards.Hand) (cards.Hand, error) {
	var guessedHand = cards.Hand{}
	fmt.Println(`
Enter your guess one card at a time, from left to right. 
Enter each card in format #S where # is the numerical card value (1-13), and S is the first letter of the suit. For example:
	5H = 5 of Hearts
	1C = Ace of Clubs
	11D = Jack of Diamonds
	12S = Queen of Spades
	13H = King of Hearts
	`)

	for i := range hand {
		promptText := fmt.Sprintf("Enter card %d: ", i+1)
		prompt := menus.NewPrompt(promptText, 13)
		guessedVal, guessedSuit, err := menus.GetUserCardInput(prompt)
		if err != nil {
			return nil, err
		}
		guessedCard := cards.NewCard(*guessedSuit, guessedVal)
		guessedHand = append(guessedHand, guessedCard)
	}

	return guessedHand, nil
}
