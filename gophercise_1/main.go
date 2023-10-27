package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	quizFileName := flag.String("q", "problems.csv", "quiz file name")
	timeToSolve := flag.Int("t", 30, "time to solve")
	flag.Parse()

	quizFile, err := os.Open(*quizFileName)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(quizFile)

	questions, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	nCorrectAnswers := 0

	fmt.Println("Press enter to start the quiz!")
	fmt.Scanln()

	timer := time.NewTimer(time.Duration(*timeToSolve) * time.Second)

	for _, qAndA := range questions {
		q, a := qAndA[0], qAndA[1]

		fmt.Println(q)
		userAnswer := make(chan string)
		go func() {
			answer := ""
			fmt.Scanln(&answer)
			userAnswer <- answer
		}()

		select {
		case answer := <-userAnswer:
			if answer == a {
				nCorrectAnswers++
			}
		case <-timer.C:
			fmt.Printf("\nYou solved %d of %d correct\n", nCorrectAnswers, len(questions))
			return
		}
	}

	fmt.Printf("You solved %d of %d correct\n", nCorrectAnswers, len(questions))
}
