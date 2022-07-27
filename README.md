# golang_quiz_project
Simple Golang quiz design, that loads questions from csv file to ask user. User can also provide a file of their own content to the program. At the end of the quiz, users' will see their quiz scores


## Features
- Users can provide path to their own quiz file as a flag when running the package
- Users can see their scores and failed at the end of the quiz
- New Feature: Quiz uses a timer, and users can set timer to start quiz


### Usage
Running by default using default questions
```shell
go run main.go
```

Running with your own questions
```shell
go run main.go -f "Path to your valid questions csv file"
```


