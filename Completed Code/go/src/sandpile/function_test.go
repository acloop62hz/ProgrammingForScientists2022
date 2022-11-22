package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

type Inputs struct {
	board Board
	x, y  int
	queue Queue
}

type Answer struct {
	board Board
	sides []int
}

func ReadFilesFromDirectory(directory string) []os.FileInfo {
	dirContents, err := ioutil.ReadDir(directory)
	if err != nil {
		panic("Error reading directory: " + directory)
	}

	return dirContents
}

func TestTopple(t *testing.T) {

	type test struct {
		input  Inputs
		answer Answer
	}

	inputDirectory := "tests/Topple/input/"
	outputDirectory := "tests/Topple/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//we now will need to create our array of tests
	tests := make([]test, len(inputFiles))

	//first, range through the input and output files and set the test values
	for i := range inputFiles {
		tests[i].input = ReadInputsFromFile(inputDirectory, inputFiles[i])
		tests[i].answer = ReadAnswerFromFile(outputDirectory, outputFiles[i])
	}

	for i, test := range tests {
		outcomeSides := tests[i].input.board.Topple(tests[i].input.x, tests[i].input.y, &tests[i].input.queue)
		outcomeBoard := tests[i].input.board
		CompareResults(outcomeSides, outcomeBoard, test.answer, t)
	}
}

func TestGenerateAbove4Queue(t *testing.T) {
	type test struct {
		input  Board
		answer Queue
	}

	inputDirectory := "tests/GenerateAbove4Queue/input/"
	outputDirectory := "tests/GenerateAbove4Queue/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//we now will need to create our array of tests
	tests := make([]test, len(inputFiles))

	//first, range through the input and output files and set the test values
	for i := range inputFiles {
		tests[i].input = ReadBoardFromFile(inputDirectory, inputFiles[i])
		tests[i].answer = ReadQueueFromFile(outputDirectory, outputFiles[i])
	}

	for i, test := range tests {
		outcome := GenerateAbove4Queue(&(tests[i].input))
		CompareResults2(outcome, test.answer, t)
	}

}

func CompareResults2(outcome Queue, answer Queue, t *testing.T) {
	fmt.Println("Comparing queues")
	if len(outcome) != len(answer) {
		t.Errorf("Error!")
	}
	n := len(outcome)
	for i := 0; i < n; i++ {
		fmt.Println("Index: ", i, " Your Answer: ", outcome[i], " Correct Answer: ", answer[i])
		if outcome[i] != answer[i] {
			t.Errorf("Error!")
		}
	}
}

func ReadQueueFromFile(directory string, File os.FileInfo) Queue {
	fileName := File.Name() //grab file name

	//now, read in the output file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	fieldIndex := 1
	var nqueue int
	var queue Queue
	//first, read lines and split along blank space
	Lines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	for _, Line := range Lines {
		if Line == "-" {
			fieldIndex += 1
			continue
		}

		if fieldIndex == 1 {
			currentLine := strings.Split(Line, " ")
			nqueue, _ = strconv.Atoi(currentLine[0])
		}

		if fieldIndex == 2 {
			currentLine := strings.Split(Line, " ")
			queue = make(Queue, nqueue)
			for i := range queue {
				queue[i].x, _ = strconv.Atoi(currentLine[i])
				queue[i].y, _ = strconv.Atoi(currentLine[i+nqueue])
			}
		}

	}
	return queue

}

func ReadBoardFromFile(directory string, File os.FileInfo) Board {
	fileName := File.Name() //grab file name

	//now, read in the output file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	board := make(Board, 3)
	for i := range board {
		board[i] = make([]Square, 3)
	}
	boardCount := 0

	Lines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")
	for _, Line := range Lines {
		currentLine := strings.Split(Line, " ")
		for i := 0; i < 3; i++ {
			board[boardCount][i].coinNumber, _ = strconv.Atoi(currentLine[i])
			inq, _ := strconv.Atoi(currentLine[i+3])
			if inq == 0 {
				board[boardCount][i].inqueue = false
			} else {
				board[boardCount][i].inqueue = true
			}
		}
		boardCount += 1
		continue
	}

	return board
}

func CompareResults(outcomeSides []int, outcomeBoard Board, answer Answer, t *testing.T) {
	fmt.Println("Comparing Boards")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("Square", i, j)
			fmt.Println("Correct Answer", answer.board[i][j])
			fmt.Println("Your Answer", outcomeBoard[i][j])
			if answer.board[i][j].coinNumber != outcomeBoard[i][j].coinNumber {
				t.Errorf("Error!")
			}
		}
	}
	fmt.Println("Comparing Falling sides")
	fmt.Print("top: Your Answer: ", outcomeSides[0], " Correct Answer ", answer.sides[0])
	fmt.Print("bottom: Your Answer: ", outcomeSides[1], " Correct Answer ", answer.sides[1])
	if outcomeSides[0] != answer.sides[0] || outcomeSides[1] != answer.sides[1] {
		t.Errorf("Error!")
	}
}

func ReadAnswerFromFile(directory string, outputFile os.FileInfo) Answer {
	fileName := outputFile.Name() //grab file name

	//now, read in the output file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	var answer Answer
	answer.sides = make([]int, 2)

	answer.board = make(Board, 3)
	for i := range answer.board {
		answer.board[i] = make([]Square, 3)
	}

	boardCount := 0
	fieldIndex := 1
	//first, read lines and split along blank space
	outputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	for _, outputLine := range outputLines {

		if outputLine == "-" {
			fieldIndex += 1
			continue
		}

		if fieldIndex == 1 {
			currentLine := strings.Split(outputLine, " ")
			for i := 0; i < 3; i++ {
				answer.board[boardCount][i].coinNumber, _ = strconv.Atoi(currentLine[i])
				inq, _ := strconv.Atoi(currentLine[i+3])
				if inq == 0 {
					answer.board[boardCount][i].inqueue = false
				} else {
					answer.board[boardCount][i].inqueue = true
				}
			}
			boardCount += 1
			continue
		}

		if fieldIndex == 2 {
			currentLine := strings.Split(outputLine, " ")
			answer.sides[0], _ = strconv.Atoi(currentLine[0])
			answer.sides[1], _ = strconv.Atoi(currentLine[1])
		}

	}

	return answer

}

func ReadInputsFromFile(directory string, inputFile os.FileInfo) Inputs {
	fileName := inputFile.Name() //grab file name

	//now, read in the input file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	var inputs Inputs
	var nqueue int

	inputs.board = make(Board, 3)
	for i := range inputs.board {
		inputs.board[i] = make([]Square, 3)
	}

	fieldIndex := 1
	boardCount := 0
	//first, read lines and split along blank space
	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	for _, inputLine := range inputLines {

		if inputLine == "-" {
			fieldIndex += 1
			boardCount = 0
			continue
		}

		if fieldIndex == 1 {
			currentLine := strings.Split(inputLine, " ")
			for i := 0; i < 3; i++ {
				inputs.board[boardCount][i].coinNumber, _ = strconv.Atoi(currentLine[i])
				inq, _ := strconv.Atoi(currentLine[i+3])
				if inq == 0 {
					inputs.board[boardCount][i].inqueue = false
				} else {
					inputs.board[boardCount][i].inqueue = true
				}
			}
			boardCount += 1
			continue
		}

		if fieldIndex == 2 {
			currentLine := strings.Split(inputLine, " ")
			inputs.queue = make(Queue, nqueue)
			for i := 0; i < nqueue; i++ {
				inputs.queue[i].x, _ = strconv.Atoi(currentLine[i])
				inputs.queue[i].y, _ = strconv.Atoi(currentLine[i+nqueue])
			}
		}

		if fieldIndex == 3 {
			currentLine := strings.Split(inputLine, " ")
			inputs.x, _ = strconv.Atoi(currentLine[0])
			inputs.y, _ = strconv.Atoi(currentLine[1])
			nqueue, _ = strconv.Atoi(currentLine[2])
		}
	}

	return inputs
}
