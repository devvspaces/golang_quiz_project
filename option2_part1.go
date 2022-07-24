package main

// This is another way to load the questions, using a more
// object oriented design (structs)
// Reason for not using this method: Because you how would have to loop
// twice. First for loading question in to slice of structs and other
// would be to loop through the array to ask questions.
// But Golang is very fast, depending on the use case this method could
// present faster user experience

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type QA struct {
	Question string
	Answer   string
}

func createQuestionAnswerList(data [][]string) []QA {
	var qList []QA
	for i, line := range data {
		if i > 0 { // omit header line
			var rec QA
			rec.Question = line[0]
			rec.Answer = line[1]
			qList = append(qList, rec)
		}
	}
	return qList
}

func main() {
	// open file
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// convert records to array of structs
	qaList := createQuestionAnswerList(data)

	// print the array
	fmt.Printf("%+v\n", qaList)
}
