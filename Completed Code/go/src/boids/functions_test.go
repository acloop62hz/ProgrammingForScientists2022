package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestUpV(t *testing.T) {
	type test struct {
		b          Boid
		time, maxs float64
		answer     OrderedPair
	}
	inputDirectory := "tests/UpV/input/"
	outputDirectory := "tests/UpV/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))

	//we now will need to create our array of tests
	tests := make([]test, len(inputFiles))

	for i := range inputFiles {
		tests[i].b, tests[i].time, tests[i].maxs = ReadBoidsFromFile(inputDirectory, inputFiles[i])
		tests[i].answer = ReadOrderPairFromFile(outputDirectory, outputFiles[i])
	}

	for i, test := range tests {
		outcome := UpdateVelocity(tests[i].b, tests[i].time, tests[i].maxs)

		if roundFloat(outcome.x, 2) != roundFloat(test.answer.x, 2) || roundFloat(outcome.y, 2) != roundFloat(test.answer.y, 2) {
			t.Errorf("Error! For input test dataset %d, your code gives %f,%f, and the correct velocity is %f,%f", i, outcome.x, outcome.y, test.answer.x, test.answer.y)
		} else {
			fmt.Println("Correct!")
		}
	}
}

func TestUpP(t *testing.T) {
	type test struct {
		b           Boid
		time, width float64
		answer      OrderedPair
	}
	inputDirectory := "tests/UpP/input/"
	outputDirectory := "tests/UpP/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))

	//we now will need to create our array of tests
	tests := make([]test, len(inputFiles))

	for i := range inputFiles {
		tests[i].b, tests[i].time, tests[i].width = ReadBoidsFromFile(inputDirectory, inputFiles[i])
		tests[i].answer = ReadOrderPairFromFile(outputDirectory, outputFiles[i])
	}

	for i, test := range tests {
		outcome := UpdatePosition(tests[i].b, tests[i].time, tests[i].width)

		if roundFloat(outcome.x, 2) != roundFloat(test.answer.x, 2) || roundFloat(outcome.y, 2) != roundFloat(test.answer.y, 2) {
			t.Errorf("Error! For input test dataset %d, your code gives %f,%f, and the correct velocity is %f,%f", i, outcome.x, outcome.y, test.answer.x, test.answer.y)
		} else {
			fmt.Println("Correct!")
		}
	}
}

func ReadBoidsFromFile(directory string, inputFile os.FileInfo) (Boid, float64, float64) {
	fileName := inputFile.Name() //grab file name

	//now, read in the input file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//first, read lines and split along blank space
	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	currentVals := make([]float64, 8)
	//each line of the file corresponds to a single line of the frequency map
	for _, inputLine := range inputLines {
		//read out the current line
		currentLine := strings.Split(inputLine, " ")
		for i := 0; i < 8; i++ {
			currentVals[i], err = strconv.ParseFloat(currentLine[i], 64)
			if err != nil {
				panic(err)
			}
		}
	}

	//make the Boids
	var b Boid
	var time, maxs float64
	b.position.x = currentVals[0]
	b.position.y = currentVals[1]
	b.velocity.x = currentVals[2]
	b.velocity.y = currentVals[3]
	b.acceleration.x = currentVals[4]
	b.acceleration.y = currentVals[5]
	time = currentVals[6]
	maxs = currentVals[7]
	return b, time, maxs
}

func ReadOrderPairFromFile(directory string, file os.FileInfo) OrderedPair {
	fileName := file.Name() //grab file name

	//now, read out the file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}
	var p OrderedPair
	//trim out extra space
	outputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")
	output := strings.Split(outputLines[0], " ")
	p.x, err = strconv.ParseFloat(output[0], 64)
	if err != nil {
		panic(err)
	}
	p.y, err = strconv.ParseFloat(output[1], 64)
	if err != nil {
		panic(err)
	}
	return p

}

func ReadFilesFromDirectory(directory string) []os.FileInfo {
	dirContents, err := ioutil.ReadDir(directory)
	if err != nil {
		panic("Error reading directory: " + directory)
	}

	return dirContents
}

func AssertEqualAndNonzero(length0, length1 int) {
	if length0 == 0 {
		panic("No files present in input directory.")
	}
	if length1 == 0 {
		panic("No files present in output directory.")
	}
	if length0 != length1 {
		panic("Number of files in directories doesn't match.")
	}
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
