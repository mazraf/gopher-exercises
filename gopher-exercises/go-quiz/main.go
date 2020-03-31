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
	//Make a csv flag to give manual csv file, then parse it.
	csvFlagName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 2, "Time limit in seconds")
	flag.Parse()

	file, err := os.Open(*csvFlagName)
	if err != nil {
		log.Fatalf("Error while opening the file : %s\n", *csvFlagName)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Error while reading file :%s", err)
	}
	problems := parseFile(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s =", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou have scored %d out of %d", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}

		}

	}
	fmt.Printf("You have scored %d out of %d", correct, len(problems))
}

type problem struct {
	q string
	a string
}

func parseFile(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}
