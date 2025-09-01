package questions

import (
	"errors"
	"fmt"
	"slices"
	"strconv"

	"example.com/questions-and-truth/cards"
)

type Question struct {
	ID      int
	Text    string
	Prompts []Prompt
}

type Prompt struct {
	Text     string
	MaxValue int
}

func NewQuestion(id int, text string, prompts []Prompt) Question {
	return Question{
		ID:      id,
		Text:    text,
		Prompts: prompts,
	}
}

func InitializeQuestions(deck cards.Deck) []Question {
	suitPromptText := ""
	for _, suit := range deck.Suits {
		suitPromptText += fmt.Sprintf("%d. %s\n", suit.ID, suit.Name)
	}
	suitPromptText += "Select a suit (1-4): "
	suitPrompt := []Prompt{{Text: suitPromptText, MaxValue: 4}}

	valuePromptText := "Select a card value (1-13): "
	valuePrompt := []Prompt{{Text: valuePromptText, MaxValue: 13}}

	positionPrompt := Prompt{Text: "Select a position (1-8): ", MaxValue: 8}
	positionPrompts := []Prompt{positionPrompt, positionPrompt, positionPrompt}

	return []Question{
		NewQuestion(1, "What is the sum of cards in three specified positions (1-8)?", positionPrompts),
		NewQuestion(2, "What is the total sum of cards of the specified suit?", suitPrompt),
		NewQuestion(3, "What is the total sum of face cards (Aces, Jacks, Queens, Kings)?", nil),
		NewQuestion(4, "What is the total sum of the number cards (2 - 10)?", nil),
		NewQuestion(5, "What is the total count of face cards (Aces, Jacks, Queens, Kings)?", nil),
		NewQuestion(6, "What is the total count of the number cards (2 - 10)?", nil),
		NewQuestion(7, "What is the total count of cards with the specified value (1-13)?", valuePrompt),
		NewQuestion(8, "Which positions contain cards of the specified suit?", suitPrompt),
		NewQuestion(9, "Which positions contain cards with the same value?", nil),
		NewQuestion(10, "Which positions contain cards with consecutive values?", nil),
		NewQuestion(11, "Which position(s) contains the card(s) with the highest value?", nil),
		NewQuestion(12, "Which position(s) contains the card(s) with the lowest value?", nil),
	}
}

func DisplayQuestionMenu(questions []Question) (*Answer, error) {
	var menuText string
	for _, q := range questions {
		menuText += fmt.Sprintf("%d. %s\n", q.ID, q.Text)
	}
	menuText += fmt.Sprintf("Select a question (1-%d): ", len(questions))

	selection, err := GetUserInputAsInt(menuText, len(questions)+1)
	if err != nil {
		return nil, err
	}

	selectedQuestion := questions[selection-1]

	return (&selectedQuestion).DisplayAndPromptIfNeeded()
}

func (q *Question) DisplayAndPromptIfNeeded() (*Answer, error) {
	fmt.Println("\nSelected question: ", q.Text)

	// If there are prompt(s), collect user input
	userMenuSelections := make([]int, len(q.Prompts))
	for i, prompt := range q.Prompts {
		selection, err := GetUserInputAsInt(prompt.Text, prompt.MaxValue)
		if err != nil {
			return nil, err
		}

		// check for duplicate responses when multiple prompts for a single question
		if slices.Contains(userMenuSelections, selection) {
			return nil, errors.New("value was already selected, no duplicates allowed")
		}

		userMenuSelections[i] = selection
	}
	return &Answer{
		Question:            *q,
		UserSpecifiedValues: userMenuSelections,
	}, nil
}

func GetUserInputAsInt(prompt string, maxVal int) (int, error) {
	// display prompt
	fmt.Print(prompt)

	// get user selection as integer
	var selectionStr string
	fmt.Scan(&selectionStr)
	selection, err := strconv.Atoi(selectionStr)

	// verify valid selection
	if err != nil {
		return 0, err
	}
	if selection < 1 || selection > maxVal {
		errMessage := fmt.Sprintf("invalid selection, must be an integer between 1 and %d", maxVal)
		return 0, errors.New(errMessage)
	}
	return selection, nil
}
