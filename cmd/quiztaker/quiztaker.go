package main

import (
	"fmt"

	"github.com/rajch/quizzer/pkg/quizcore"
)

func processQuestion(q quizcore.Simplequestion) bool {

	fmt.Println(q.Question())
	fmt.Print("Answer? ")

	var ans string
	_, err := fmt.Scanln(&ans)
	if err != nil {
		return false
	}

	return q.CheckAnswer(ans)
}

func processQuiz(questions []quizcore.Simplequestion) {
	numquestions := len(questions)
	if numquestions < 1 {
		return
	}

	fmt.Printf("Welcome to the quiz. There are %d questions.\n", numquestions)

	score := 0

	for i, q := range questions {
		fmt.Println("*****************************************")
		fmt.Printf("Score: %d\n", score)
		fmt.Printf("Question %d/%d:\n", i+1, numquestions)
		if processQuestion(q) {
			fmt.Println("CORRECT!")
			score++
		} else {
			fmt.Println("WRONG!")
		}
	}

	fmt.Println("*****************************************")
	fmt.Printf("\n\nFinal Score: %d\n", score)
}
