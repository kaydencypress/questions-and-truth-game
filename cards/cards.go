package cards

import (
	"errors"
	"math/rand"
)

var GameDeck = NewDeck()

type Deck struct {
	AvailableCards []Card
	Suits          []Suit
}

type Suit struct {
	ID   int
	Name string
}

type Card struct {
	Suit       Suit
	Value      int
	IsFaceCard bool
	FaceValue  string
}

func NewCard(suit Suit, val int) Card {
	var faceValueMap = map[int]string{
		1:  "Ace",
		11: "Jack",
		12: "Queen",
		13: "King",
	}

	faceVal, isFaceCard := faceValueMap[val]

	return Card{
		Suit:       suit,
		Value:      val,
		IsFaceCard: isFaceCard,
		FaceValue:  faceVal,
	}
}

func NewDeck() Deck {
	suits := []Suit{
		{ID: 1, Name: "Clubs"},
		{ID: 2, Name: "Diamonds"},
		{ID: 3, Name: "Hearts"},
		{ID: 4, Name: "Spades"},
	}
	var cards = make([]Card, 0, 52)

	for _, suit := range suits {
		for val := 1; val <= 13; val++ {
			cards = append(cards, NewCard(suit, val))
		}
	}
	return Deck{
		AvailableCards: cards,
		Suits:          suits,
	}
}

func (deck *Deck) ShuffleDeck() {
	rand.Shuffle(len(deck.AvailableCards), func(i, j int) {
		deck.AvailableCards[i], deck.AvailableCards[j] = deck.AvailableCards[j], deck.AvailableCards[i]
	})
}

func (deck *Deck) PickCard() (*Card, error) {
	if len(deck.AvailableCards) < 1 {
		return nil, errors.New("not enough cards left in the deck")
	}

	card := deck.AvailableCards[0]
	deck.AvailableCards = deck.AvailableCards[1:]

	return &card, nil
}
