package questions

import (
	"errors"
	"fmt"
	"strconv"
)

type Question struct {
	Category string
	Text     string
	Prompts  []Prompt
}

type Prompt struct {
	Text     string
	MaxValue int
}

func NewQuestion(category string, text string, prompts []Prompt) Question {
	return Question{
		Category: category,
		Text:     text,
		Prompts:  prompts,
	}
}

func (q *Question) DisplayAndPromptIfNeeded() (*Answer, error) {
	fmt.Println(q.Text)

	// If there are prompt(s), collect user selection(s) as integer(s)
	userMenuSelections := make([]int, len(q.Prompts))
	for i, prompt := range q.Prompts {
		// display prompt
		fmt.Print(prompt.Text)

		// get user selection as integer
		var selectionStr string
		fmt.Scan(&selectionStr)
		selection, err := strconv.Atoi(selectionStr)

		// verify valid selection
		if err != nil {
			return nil, err
		}
		if selection < 1 || selection > prompt.MaxValue {
			errMessage := fmt.Sprintf("invalid selection, must be an integer between 1 and %d", prompt.MaxValue)
			return nil, errors.New(errMessage)
		}

		// store user selection
		userMenuSelections[i] = selection
	}
	return &Answer{
		Question:            *q,
		UserSpecifiedValues: userMenuSelections,
	}, nil
}
