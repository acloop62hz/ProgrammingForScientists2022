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

func TestListMersennePrimes(t *testing.T) {
	type test struct {
		n      int
		output []int
	}
	inputDirectory := "tests/MersennePrimes/input/"
	outputDirectory := "tests/MersennePrimes/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))

	//we now will need to create our array of tests
	tests := make([]test, len(inputFiles))

	//first, range through the input and output files and set the test values
	for i := range inputFiles {
		tests[i].n = ReadIntegerFromFile(inputDirectory, inputFiles[i])
		tests[i].output = ReadIntArrayFromFile(outputDirectory, outputFiles[i])
	}

	//are the tests correct?
	for i, test := range tests {
		outcome := ListMersennePrimes(test.n)
		flag := true
		for j := range outcome {
			if outcome[j] != test.output[j] {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println("Correct! When the intput is", test.n, "the primers are", test.output)
		} else {
			t.Errorf("Error! For input test dataset %d, your code gives %v, and the correct answer is %v", i, outcome, test.output)
		}
	}

}

func ReadIntegerFromFile(directory string, file os.FileInfo) int {
	//now, consult the associated output file.
	fileName := file.Name() //grab file name

	//now, read out the file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//trim out extra space and store as a slice of strings, each containing one line.
	outputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	//parse the float
	answer, err := strconv.Atoi(outputLines[0])

	if err != nil {
		panic(err)
	}

	return answer
}

func ReadIntArrayFromFile(directory string, file os.FileInfo) []int {
	fileName := file.Name() //grab file name

	//now, read out the file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//trim out extra space
	outputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")
	output := strings.Split(outputLines[0], " ")
	intOutput := make([]int, len(output))
	for i := range output {
		intOutput[i], err = strconv.Atoi(output[i])
		if err != nil {
			panic(err)
		}
	}
	return intOutput

}

func ListMersennePrimes(n int) []int {
	list := []int{}
	for i := 1; i <= n; i++ {
		r := int(math.Pow(2, float64(i)) - 1)
		if IsPrime(r) {
			list = append(list, r)
		}
	}
	return list
}

func IsPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
