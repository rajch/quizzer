package quizcore

// Simplequestion with one answer
type Simplequestion struct {
	question string
	answer   string
}

// Question to be asked
func (q Simplequestion) Question() string {
	return q.question
}

// Answer to the question
func (q Simplequestion) Answer() string {
	return q.answer
}

// CheckAnswer returns true if answer is correct, false otherwise
func (q Simplequestion) CheckAnswer(answer string) bool {
	return q.Answer() == answer
}

// NewSimpleQuestion creates a new simple question
func NewSimpleQuestion(newquestion string, newanswer string) Simplequestion {
	return Simplequestion{
		question: newquestion,
		answer:   newanswer,
	}
}
