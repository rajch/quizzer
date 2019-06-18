package quizcore

// Simplequestion with one answer
type Simplequestion struct {
	question string
	answer   string
}

// Question to be asked
func (q *Simplequestion) Question() string {
	return q.question
}

// Answer to the question
func (q *Simplequestion) Answer() string {
	return q.question
}

// NewSimpleQuestion creates a new simple question
func NewSimpleQuestion(newquestion string, newanswer string) Simplequestion {
	return Simplequestion{
		question: newquestion,
		answer:   newanswer,
	}
}
