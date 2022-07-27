package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type QA struct {
	q string
	a string
}

func parseLines(data [][]string) []QA {
	problems := make([]QA, len(data))
	for index, line := range data {
		p := &problems[index]
		p.q = line[0]
		p.a = normalizeString(line[1])
	}
	return problems
}

func normalizeString(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func exitErr(value string) {
	fmt.Println(value)
	os.Exit(1)
}

func startQuiz(problems []QA) {
	line_reader := bufio.NewReader(os.Stdin)
	correct := 0

	for index, problem := range problems {
		fmt.Printf("Question %d: %s -> ", index+1, problem.q)
		answer, _ := line_reader.ReadString('\n')
		answer = normalizeString(answer)
		if answer == problem.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func getOptions() (string, bool) {
	csvFileName := flag.String("c", "problems.csv", "Filename for the csv containing questions")
	shuffle := flag.Bool("s", false, "Set to 1 to shuffle the questions")
	flag.Parse()
	return *csvFileName, *shuffle
}

func readFile(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		value := fmt.Sprintf("Error trying to open file: %s", name)
		exitErr(value)
	}
	return file
}

func getFileLines(name string) [][]string {
	file := readFile(name)
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		value := fmt.Sprintf("Error trying to read file: %s", name)
		exitErr(value)
	}

	return lines
}

func randomShuffle(problems *[]QA) {
	for i := len(*problems) - 1; i > 0; i-- {
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(i + 1)
		(*problems)[i], (*problems)[j] = (*problems)[j], (*problems)[i]
	}
}

func main() {

	name, shuffle := getOptions()
	lines := getFileLines(name)
	problems := parseLines(lines)
	if shuffle {
		randomShuffle(&problems)
	}
	startQuiz(problems)

}
