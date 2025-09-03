package menus

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
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
