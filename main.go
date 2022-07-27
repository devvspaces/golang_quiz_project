package main

import (
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

func startQuiz(problems []QA, ch chan bool) {
	fmt.Println("Quiz has started")

	correct := 0

	go func() {
		for index, problem := range problems {
			fmt.Printf("Question %d: %s -> ", index+1, problem.q)
			var answer string
			fmt.Scanf("%s\n", &answer)
			answer = normalizeString(answer)
			if answer == problem.a {
				correct++
			}
		}
		ch <- true
	}()

	<-ch
	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
}

func getOptions() (string, bool, int) {
	csvFileName := flag.String("c", "problems.csv", "Filename for the csv containing questions")
	shuffle := flag.Bool("s", false, "Set to 1 to shuffle the questions")
	duration := flag.Int("d", 30, "Duration for quiz value in seconds")
	flag.Parse()
	return *csvFileName, *shuffle, *duration
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

func timer(ch chan bool, duration int) {
	time.Sleep(time.Second * time.Duration(duration))
	ch <- true
}

func main() {

	ch := make(chan bool)

	name, shuffle, duration := getOptions()
	lines := getFileLines(name)
	problems := parseLines(lines)
	if shuffle {
		randomShuffle(&problems)
	}
	fmt.Print("Press Enter to start your quiz: ")
	fmt.Scanf("%s\n")
	go timer(ch, duration)
	startQuiz(problems, ch)

}
