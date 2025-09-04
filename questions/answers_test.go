package questions

import (
	"testing"

	"example.com/questions-and-truth/cards"
)

var handWithMatches = cards.Hand{
	cards.NewCard(cards.Clubs, 2),
	cards.NewCard(cards.Spades, 1),
	cards.NewCard(cards.Spades, 9),
	cards.NewCard(cards.Diamonds, 10),
	cards.NewCard(cards.Clubs, 7),
	cards.NewCard(cards.Diamonds, 11),
	cards.NewCard(cards.Hearts, 12),
	cards.NewCard(cards.Spades, 11),
}

var handNoMatches = cards.Hand{
	cards.NewCard(cards.Spades, 2),
	cards.NewCard(cards.Spades, 4),
	cards.NewCard(cards.Diamonds, 6),
	cards.NewCard(cards.Diamonds, 8),
	cards.NewCard(cards.Clubs, 7),
	cards.NewCard(cards.Spades, 6),
	cards.NewCard(cards.Diamonds, 10),
	cards.NewCard(cards.Clubs, 9),
}

var handOnlyFaceCards = cards.Hand{
	cards.NewCard(cards.Clubs, 1),
	cards.NewCard(cards.Diamonds, 1),
	cards.NewCard(cards.Clubs, 11),
	cards.NewCard(cards.Spades, 1),
	cards.NewCard(cards.Diamonds, 12),
	cards.NewCard(cards.Diamonds, 13),
	cards.NewCard(cards.Spades, 12),
	cards.NewCard(cards.Hearts, 1),
}

type answerTest struct {
	name   string
	want   string
	hand   cards.Hand
	inputs []int
}

func TestGetSumOfCardsInPositions(t *testing.T) {
	var tests = []answerTest{
		{
			name:   "GetSumOfCardsInPositions-matches",
			want:   "25",
			hand:   handWithMatches,
			inputs: []int{1, 6, 7},
		},
	}
	runTests(t, tests, GetSumOfCardsInPositions)
}

func TestGetSumOfSuits(t *testing.T) {
	var tests = []answerTest{
		{
			name:   "GetSumOfSuits-matches",
			want:   "21",
			hand:   handWithMatches,
			inputs: []int{cards.Spades.ID},
		},
		{
			name:   "GetSumOfSuits-no-matches",
			want:   "0",
			hand:   handNoMatches,
			inputs: []int{cards.Hearts.ID},
		},
	}
	runTests(t, tests, GetSumOfSuits)
}

func TestGetSumOfFaceCards(t *testing.T) {
	var tests = []answerTest{
		{
			name:   "GetSumOfFaceCards-matches",
			want:   "35",
			hand:   handWithMatches,
			inputs: []int{},
		},
		{
			name:   "GetSumOfFaceCards-no-matches",
			want:   "0",
			hand:   handNoMatches,
			inputs: []int{},
		},
	}
	runTests(t, tests, GetSumOfFaceCards)
}

func TestGetSumOfNumberCards(t *testing.T) {
	var tests = []answerTest{
		{
			name:   "GetSumOfNumberCards-matches",
			want:   "28",
			hand:   handWithMatches,
			inputs: []int{},
		},
		{
			name:   "GetSumOfNumberCards-no-matches",
			want:   "0",
			hand:   handOnlyFaceCards,
			inputs: []int{},
		},
	}
	runTests(t, tests, GetSumOfNumberCards)
}

func runTests(t *testing.T, tests []answerTest, testFunc func(cards.Hand, []int) string) {
	for _, test := range tests {
		actualVal := testFunc(test.hand, test.inputs)
		t.Run(test.name, func(t *testing.T) {
			if actualVal != test.want {
				t.Errorf(`Result is incorrect, got %s, wanted %s`, actualVal, test.want)
			}
		})
	}
}
