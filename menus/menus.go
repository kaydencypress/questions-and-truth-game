package menus

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"example.com/questions-and-truth/cards"
)

type Prompt struct {
	Text     string
	MaxValue int
}

func NewPrompt(text string, maxValue int) Prompt {
	return Prompt{
		Text:     text,
		MaxValue: maxValue,
	}
}

func GetUserInputAsInt(prompt Prompt) (int, error) {
	// display prompt
	fmt.Print(prompt.Text)

	// get user selection as integer
	var selectionStr string
	fmt.Scan(&selectionStr)
	selection, err := strconv.Atoi(selectionStr)

	// verify valid selection
	if err != nil {
		return 0, err
	}
	if selection < 1 || selection > prompt.MaxValue {
		errMessage := fmt.Sprintf("invalid selection, must be an integer between 1 and %d", prompt.MaxValue)
		return 0, errors.New(errMessage)
	}
	return selection, nil
}

func GetAllUserInputsAsInt(prompts []Prompt) ([]int, error) {
	// If there are prompt(s), collect user input
	userMenuSelections := make([]int, len(prompts))
	for i, prompt := range prompts {
		selection, err := GetUserInputAsInt(prompt)
		if err != nil {
			return nil, err
		}

		// check for duplicate responses when multiple prompts for a single question
		if slices.Contains(userMenuSelections, selection) {
			return nil, errors.New("value was already selected, no duplicates allowed")
		}

		userMenuSelections[i] = selection
	}
	return userMenuSelections, nil
}

func GetUserCardInput(prompt Prompt) (int, *cards.Suit, error) {
	// display prompt
	fmt.Print(prompt.Text)

	// get user selection as string
	var selectionStr string
	fmt.Scan(&selectionStr)

	if len(selectionStr) < 2 {
		return 0, nil, errors.New("invalid input, not enough characters")
	}

	// parse and validate suit and card value
	suitStr := selectionStr[len(selectionStr)-1:]
	suitStr = strings.ToUpper(suitStr)
	suit, err := parseSuitInput(suitStr)
	if err != nil {
		return 0, nil, err
	}

	cardValStr := selectionStr[:len(selectionStr)-1]
	cardVal, err := parseIntInput(cardValStr, prompt.MaxValue)
	if err != nil {
		return 0, nil, err
	}

	return cardVal, suit, nil
}

func parseIntInput(input string, maxVal int) (int, error) {
	// parse integer
	inputInt, err := strconv.Atoi(input)

	if err != nil {
		return 0, err
	}

	// validate input is in valid range
	if inputInt < 1 || inputInt > maxVal {
		errMessage := fmt.Sprintf("invalid input, must be an integer between 1 and %d", maxVal)
		return 0, errors.New(errMessage)
	}
	return inputInt, nil
}

func parseSuitInput(input string) (*cards.Suit, error) {
	// validate input is a single character
	if len(input) != 1 {
		return nil, errors.New("invalid input for suit, must be a single character")
	}

	// validate input matches the first character of a valid suit
	for _, suit := range cards.GameDeck.Suits {
		if input == suit.Name[0:1] {
			return &suit, nil
		}
	}

	return nil, errors.New("invalid input, does not match first letter of any suit")
}

func SelectQuestionOrTruth() (int, error) {
	fmt.Printf("\nAsk a question or guess the truth\n1. Question\n2. Truth\n")
	prompt := NewPrompt("Select option (1-2): ", 2)
	return GetUserInputAsInt(prompt)
}
