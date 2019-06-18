package main

import "github.com/rajch/quizzer/pkg/quizcore"

var questions []quizcore.Simplequestion

func initquestions() {
	questions = []quizcore.Simplequestion{
		quizcore.NewSimpleQuestion(
			"What is your favorite colour? ",
			"black",
		),
		quizcore.NewSimpleQuestion(
			"What is 2+2? ",
			"5",
		),
		quizcore.NewSimpleQuestion(
			"What letter comes after j? ",
			"k",
		),
	}
}

func main() {
	initquestions()

	processQuiz(questions)
}
