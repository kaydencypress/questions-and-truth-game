package cards

import (
	"math/rand"
	"sort"
)

type Hand []Card

func NewHand(deck *Deck, n int) (Hand, error) {
	hand := make([]Card, n)
	for i := range n {
		selectedCard, err := deck.PickCard()
		if err != nil {
			return nil, err
		}
		hand[i] = *selectedCard
	}
	return hand, nil
}

func (hand *Hand) ReorderHand() {
	// cards within suit must be arranged in ascending order, so sorting each suit seperately
	cardsBySuit := make([][]Card, 4)
	for _, card := range *hand {
		// seperate cards by suit ID
		currentSuit := card.Suit
		cardsBySuit[currentSuit.ID-1] = append(cardsBySuit[currentSuit.ID-1], card)
	}

	for _, cardSet := range cardsBySuit {
		// don't need to sort if there are no cards of this suit
		if len(cardSet) == 0 {
			continue
		}

		// reorder cards in ascending value
		sort.Slice(cardSet, func(i, j int) bool {
			return cardSet[i].Value < cardSet[j].Value
		})
	}

	// merge seperate suits into newly ordered hand
	newHand := make(Hand, 0, len(*hand))
	for {
		// stop when all cards have been arranged
		if len(cardsBySuit) == 0 {
			break
		}

		// randomly select a suit
		chosenSuit := rand.Intn(len(cardsBySuit))

		// if there are no cards of this suit left, remove suit
		if len(cardsBySuit[chosenSuit]) == 0 {
			if chosenSuit == len(cardsBySuit)-1 {
				cardsBySuit = cardsBySuit[:chosenSuit]
			} else {
				cardsBySuit = append(cardsBySuit[:chosenSuit], cardsBySuit[chosenSuit+1:]...)
			}
			continue
		}

		// select card with smallest value
		newHand = append(newHand, cardsBySuit[chosenSuit][0])
		cardsBySuit[chosenSuit] = cardsBySuit[chosenSuit][1:]
	}

	*hand = newHand
}
