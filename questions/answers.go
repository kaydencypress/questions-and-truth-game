package questions

import (
	"fmt"
	"strconv"

	"example.com/questions-and-truth/cards"
)

type AnswerFunc func(hand cards.Hand, userInputs []int) string

func GetSumOfCardsInPositions(hand cards.Hand, positions []int) string {
	var sum int
	for _, pos := range positions {
		sum += hand[pos-1].Value
	}
	return strconv.Itoa(sum)
}

func GetSumOfSuits(hand cards.Hand, suitIds []int) string {
	var sum int
	for _, suitId := range suitIds {
		for _, card := range hand {
			if card.Suit.ID == suitId {
				sum += card.Value
			}
		}
	}
	return strconv.Itoa(sum)
}

func GetSumOfFaceCards(hand cards.Hand, _ []int) string {
	var sum int
	for _, card := range hand {
		if card.IsFaceCard {
			sum += card.Value
		}
	}
	return strconv.Itoa(sum)
}

func GetSumOfNumberCards(hand cards.Hand, _ []int) string {
	var sum int
	for _, card := range hand {
		if !card.IsFaceCard {
			sum += card.Value
		}
	}
	return strconv.Itoa(sum)
}

func GetCountOfFaceCards(hand cards.Hand, _ []int) string {
	var sum int
	for _, card := range hand {
		if card.IsFaceCard {
			sum += 1
		}
	}
	return strconv.Itoa(sum)
}

func GetCountOfNumberCards(hand cards.Hand, _ []int) string {
	var sum int
	for _, card := range hand {
		if !card.IsFaceCard {
			sum += 1
		}
	}
	return strconv.Itoa(sum)
}

func GetCountOfCardValues(hand cards.Hand, values []int) string {
	var sum int
	for _, val := range values {
		for _, card := range hand {
			if card.Value == val {
				sum += 1
			}
		}
	}
	return strconv.Itoa(sum)
}

func GetPositionsOfSuits(hand cards.Hand, suitIds []int) string {
	var positions = make([]int, 0, 8)
	for _, suitId := range suitIds {
		for i, card := range hand {
			if card.Suit.ID == suitId {
				positions = append(positions, i+1)
			}
		}
	}
	return formatPositionResult(positions)
}

func GetPositionsOfSameValues(hand cards.Hand, _ []int) string {
	var positions = make([]int, 0, 8)
	positionsByVal := make(map[int][]int, 8)
	for i, card := range hand {
		_, isDupVal := positionsByVal[card.Value]
		positionsByVal[card.Value] = append(positionsByVal[card.Value], i+1)
		if isDupVal {
			positions = append(positions, positionsByVal[card.Value]...)
		}
	}
	return formatPositionResult(positions)
}

func GetPositionsOfConsecutiveCards(hand cards.Hand, _ []int) string {
	positions := make([]int, 0, 8)
	previousValueBySuit := make(map[string]int, 4)

	for i, card := range hand {
		previousVal, ok := previousValueBySuit[card.Suit.Name]
		if ok {
			if card.Value-previousVal == 1 {
				positions = append(positions, i+1)
			}
		}
		previousValueBySuit[card.Suit.Name] = card.Value
	}

	return formatPositionResult(positions)
}

func GetPositionsOfHighestCards(hand cards.Hand, _ []int) string {
	positions := make([]int, 0, 4)
	var highestVal int

	for i, card := range hand {

		if card.Value == highestVal {
			positions = append(positions, i+1)
			continue
		}

		if card.Value > highestVal {
			highestVal = card.Value
			positions = []int{i + 1} // overwrite
		}
	}

	return formatPositionResult(positions)
}

func GetPositionsOfLowestCards(hand cards.Hand, _ []int) string {
	positions := make([]int, 0, 4)
	var lowestVal int

	for i, card := range hand {

		if card.Value == lowestVal {
			positions = append(positions, i+1)
			continue
		}

		if card.Value < lowestVal {
			lowestVal = card.Value
			positions = []int{i + 1} // overwrite
		}
	}
	return formatPositionResult(positions)
}

func formatPositionResult(positions []int) string {
	if len(positions) == 0 {
		return "None"
	}
	return fmt.Sprintln(positions)
}
