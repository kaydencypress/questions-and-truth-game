package questions

import "errors"

type Answer struct {
	Question            Question
	UserSpecifiedValues []int
	Response            string
}

func (answer Answer) GetAnswer() error {
	return errors.New("not implemented")
}
