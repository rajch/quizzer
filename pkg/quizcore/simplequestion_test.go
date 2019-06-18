package quizcore

import "testing"

func TestSimpleQuestion(t *testing.T) {
	question := "What is 2+2?"
	answer := "4"
	q := NewSimpleQuestion(question, answer)

	if q.question != question || q.Question() != question {
		t.Log("Problem with question.")
		t.Fail()
	}

	if q.answer != answer || q.Answer() != answer {
		t.Log("Problem with answer.")
		t.Fail()
	}

}
