package main

import (
	"encoding/csv"
	"flag"
	"os"
	"strings"
	"fmt"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in th format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("failed to open csv file: %s\n", *csvFilename)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Printf("failed to read csv file: %s\n", *csvFilename)
		os.Exit(1)
	}
	problems := parseLines(lines)	
	fmt.Println(problems)


	correct := 0
	for i, p := range problems {
		fmt.Printf("problem # %d:, %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer) // gets rid of all spaces, read stdin into answer

		if answer == p.a {
			fmt.Print("correct!\n")
			correct ++
		}
	}

	fmt.Printf("You scored %d correct out of %d questions", correct, len(problems))
} // end main

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{q: line[0], a: strings.TrimSpace(line[1])}
	}
	return ret
}

type problem struct {
	q string
	a string
}