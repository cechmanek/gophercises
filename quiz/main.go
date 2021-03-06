package main

import (
	"encoding/csv"
	"flag"
	"os"
	"strings"
	"fmt"
)

func main() {
	// generic helper flags so you can pass args like "quiz.exe -csv problems.csv"
	csvFilename := flag.String("csv", "problems.csv",
														 "a csv file in th format of 'question,answer'")
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

	// enumerate problems and ask user for answers. Also track correct responses
	correct := 0
	for i, problem := range problems {
		fmt.Printf("problem # %d:, %s = \n", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer) // gets rid of all spaces, read stdin into answer

		if answer == problem.answer {
			fmt.Print("correct!\n")
			correct ++
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