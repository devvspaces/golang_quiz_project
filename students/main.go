package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func get_file_name() string {
	defer func() {
		str := recover()
		if str != nil {
			fmt.Println("Make sure you passed a file name")
		}
	}()

	allowed_flags := []string{
		"-f",
	}

	if len(os.Args) > 1 {
		flag := os.Args[1]
		found := false
		for _, value := range allowed_flags {
			if flag == value {
				found = true
				file_name := os.Args[2]
				return file_name
			}
		}
		if found == false {
			panic("Invalid Flag used")
		}
	}

	return "problems.csv"
}

func trim(value string) string {
	return strings.Trim(value, " ")
}

func main() {
	// open file
	f, err := os.Open(get_file_name())
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Golang Mini Quiz")
	fmt.Println("---------------------")

	question_count := 1
	correct := 0
	failed := 0

	for {
		data, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		question := data[0]
		answer := trim(data[1])

		fmt.Printf("Question (%d): %s\n", question_count, question)
		fmt.Print("-> ")

		text, _ := reader.ReadString('\n')

		// convert CRLF to LF
		text = strings.Replace(text, "\r", "", -1)
		text = strings.Replace(text, "\n", "", -1)
		text = trim(text)

		question_count += 1

		if strings.Compare(answer, text) == 0 {
			correct += 1
			continue
		}

		failed += 1

	}

	fmt.Println("---------------------")
	fmt.Println("Quiz Ended")
	fmt.Println("Your Score")
	fmt.Printf("Passed: %d\n", correct)
	fmt.Printf("Failed: %d\n", failed)

}
