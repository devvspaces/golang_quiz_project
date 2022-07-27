# golang_quiz_project
Simple Golang quiz design, that loads questions from csv file to ask user. User can also provide a file of their own content to the program. At the end of the quiz, users' will see their quiz scores


## Features
- Users can provide path to their own quiz file as a flag when running the package
- Users can see their scores and failed at the end of the quiz
- New Feature: You can provide an argument to shuffle your quiz
- New Feature: Quiz uses a timer, and users can set timer to start quiz


### Usage
Checking help to see flags documentations
```shell
go run main.go -h
```

Running by default using default questions, timer and shuffle
```shell
go run main.go
```

Running with your own questions
This will load the quesions from the path you have provided
```shell
go run main.go -c "Path to your valid questions csv file"
```

Running with your own timer
Sets timer to 40 seconds
```shell
go run main.go -d 40
```

Shuffling questions
Uses flag to shuffle questions in random order
```shell
go run main.go -s
```
