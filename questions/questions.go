package questions

import (
	"fmt"

	"example.com/questions-and-truth/cards"
	"example.com/questions-and-truth/menus"
)

type Question struct {
	ID        int
	Text      string
	Prompts   []menus.Prompt
	GetAnswer AnswerFunc
}

func NewQuestion(id int, text string, prompts []menus.Prompt, getAnswer AnswerFunc) Question {
	return Question{
		ID:        id,
		Text:      text,
		Prompts:   prompts,
		GetAnswer: getAnswer,
	}
}

func InitializeQuestions() []Question {
	suitPromptText := ""
	for _, suit := range cards.GameDeck.Suits {
		suitPromptText += fmt.Sprintf("%d. %s\n", suit.ID, suit.Name)
	}
	suitPromptText += "Select a suit (1-4): "
	suitPrompt := menus.NewPrompt(suitPromptText, 4)
	suitPrompts := []menus.Prompt{suitPrompt}

	valuePrompt := menus.NewPrompt("Select a card value (1-13): ", 13)
	valuePrompts := []menus.Prompt{valuePrompt}

	positionPrompt := menus.NewPrompt("Select a position (1-8): ", 8)
	positionPrompts := []menus.Prompt{positionPrompt, positionPrompt, positionPrompt}

	return []Question{
		NewQuestion(1, "What is the sum of cards in three specified positions (1-8)?", positionPrompts, GetSumOfCardsInPositions),
		NewQuestion(2, "What is the total sum of cards of the specified suit?", suitPrompts, GetSumOfSuits),
		NewQuestion(3, "What is the total sum of face cards (Aces, Jacks, Queens, Kings)?", nil, GetSumOfFaceCards),
		NewQuestion(4, "What is the total sum of the number cards (2 - 10)?", nil, GetSumOfNumberCards),
		NewQuestion(5, "What is the total count of face cards (Aces, Jacks, Queens, Kings)?", nil, GetCountOfFaceCards),
		NewQuestion(6, "What is the total count of the number cards (2 - 10)?", nil, GetCountOfNumberCards),
		NewQuestion(7, "What is the total count of cards with the specified value (1-13)?", valuePrompts, GetCountOfCardValues),
		NewQuestion(8, "Which position(s) contain cards of the specified suit?", suitPrompts, GetPositionsOfSuits),
		NewQuestion(9, "Which position(s) contain cards with the same value?", nil, GetPositionsOfSameValues),
		NewQuestion(10, "Which position(s) contain cards that are consecutive to the previous card of the same suit?", nil, GetPositionsOfConsecutiveCards),
		NewQuestion(11, "Which position(s) contains the card(s) with the highest value?", nil, GetPositionsOfHighestCards),
		NewQuestion(12, "Which position(s) contains the card(s) with the lowest value?", nil, GetPositionsOfLowestCards),
	}

}

func SelectQuestionAndGetAnswer(questions []Question, hand cards.Hand) error {
	// prompt user to select a question
	selectedQuestion, err := SelectQuestion(questions)
	if err != nil {
		return err
	}

	// collect additional user input for question if needed
	userInputs, err := selectedQuestion.DisplayQuestionAndGetUserInputs()
	if err != nil {
		return err
	}

	// get answer and print result
	result := selectedQuestion.GetAnswer(hand, userInputs)
	fmt.Printf("\nAnswer: %v\n", result)
	return nil
}

func SelectQuestion(questions []Question) (*Question, error) {
	// print available questions
	fmt.Printf("\nAvailable questions:\n")
	for _, q := range questions {
		fmt.Printf("%d. %s\n", q.ID, q.Text)
	}

	// prompt for user's selection
	promptText := fmt.Sprintf("Select a question (1-%d): ", len(questions))
	prompt := menus.NewPrompt(promptText, len(questions))
	selection, err := menus.GetUserInputAsInt(prompt)
	if err != nil {
		return nil, err
	}

	selectedQuestion := questions[selection-1]
	return &selectedQuestion, nil
}

func (q *Question) DisplayQuestionAndGetUserInputs() ([]int, error) {
	fmt.Println("\nSelected question: ", q.Text)

	// If there are prompt(s), collect user input
	userMenuSelections, err := menus.GetAllUserInputsAsInt(q.Prompts)

	if err != nil {
		return nil, err
	}

	return userMenuSelections, nil
}
