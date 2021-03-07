package main

import (
	"encoding/csv"
	"flag"
	"os"
	"strings"
	"fmt"
	"time"
)

func main() {
	// generic helper flags so you can pass args like "quiz.exe -csv problems.csv"
	csvFilename := flag.String("csv", "problems.csv",
														 "a csv file in th format of 'question,answer'")
	timeLimit := flag.Int("limit", 5, "time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename) // csvFilename is a pointer
	if err != nil {
		fmt.Printf("failed to open csv file: %s\n", *csvFilename)
		os.Exit(1)
	}

	// attempt to read in all problems
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("failed to read csv file: %s\n", *csvFilename)
		os.Exit(1)
	}
	problems := parseLines(lines)	
	fmt.Println(problems)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// enumerate problems and ask user for answers. Also track correct responses
	correct := 0
	for i, problem := range problems {
		fmt.Printf("problem # %d:, %s = \n", i+1, problem.question)

		answerChannel := make(chan string)
		go func(){
			var answer string
			fmt.Scanf("%s\n", &answer) // strip white space, read stdin into answer
			answerChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("time is up, you scored %d out of %d", correct, len(problems))
			return // exit main() function, ending program
		case answer := <-answerChannel:
			if answer == problem.answer {
				fmt.Print("correct!\n")
				correct ++
			}
		}
	}

	fmt.Printf("You scored %d correct from %d questions", correct, len(problems))
} // end main

// helper functions and structs
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{question: line[0], answer: strings.TrimSpace(line[1])}
	}
	return ret
}

type problem struct {
	question string // question
	answer string // answer
}