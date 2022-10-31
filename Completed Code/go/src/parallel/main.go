package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

func main() {
	/*
		  //this is all nonsense
			n := 40
		  factorial1 := go Perm(1, 21)
		  factorial2 := go Perm(21,41)
		  fmt.Println(factorial1 * factorial2)
	*/

	a := []int{1, 2, 4, 523, 5, 1, 25, 51, 2, 4, 32, 44, 31, 51}
	results := MaxMultiProcs(a, 3)
	fmt.Println(results)
}

func SummingParallel() {
	//declare a slice of ints
	a := make([]int, 10000000)
	for i := range a {
		a[i] = i + 1
	}

	numProcs := runtime.NumCPU()

	start := time.Now()
	SumMultiProc(a, numProcs)
	elapsed := time.Since(start)
	log.Printf("Summing in parallel took %s", elapsed)

	start2 := time.Now()
	SumSerial(a)
	elapsed2 := time.Since(start2)
	log.Printf("Summing serially took %s", elapsed2)

}

func SumSerial(a []int) int {
	s := 0
	for _, val := range a {
		s += val
	}
	return s
}

func SumMultiProc(a []int, numProcs int) int {
	n := len(a)
	s := 0
	c := make(chan int)

	//split the array into numProcs approximately equal pieces
	for i := 0; i < numProcs; i++ {
		startIndex := i * (n / numProcs)
		endIndex := (i + 1) * (n / numProcs)
		if i < numProcs-1 {
			//normal case
			go Sum(a[startIndex:endIndex], c)
		} else { // i == numProcs - 1
			//end of the slice -- make sure you go to the very end
			go Sum(a[startIndex:], c)
		}
	}

	//get values from channel numProcs times and add them to the growing sum s
	for i := 0; i < numProcs; i++ {
		s += <-c
	}

	return s
}

func Sum(a []int, c chan int) {
	s := 0
	for _, v := range a {
		s += v
	}
	c <- s
}

func ParallelFactorial() {
	c := make(chan int)
	go PermChannel(1, 11, c)
	go PermChannel(11, 21, c)
	fmt.Println(<-c * <-c)
}

func PermChannel(k, n int, c chan int) {
	p := 1
	for i := k; i < n; i++ {
		p *= i
	}
	c <- p
}

func BasicChannels() {

	//we can force Go to use a different max number of procs
	runtime.GOMAXPROCS(1)

	//channels store a value of a given type and allow functions to communicate with each other
	c := make(chan string) // this channel is "synchronous"
	//c <- "Hello" // the channel blocks, meaning that it will not continue in this serial process until someone is on the other end of the channel ready to receive the message
	go SayHi(c)
	fmt.Println(<-c)

	fmt.Println("Exiting normally.")
}

func SayHi(c chan string) {
	fmt.Println("Yo")
	c <- "Hello!"
	//only block what comes after this in the subroutine
	//(which is nothing)
}

func Perm(k, n int) int {
	p := 1
	for i := k; i < n; i++ {
		p *= i
	}
	return p
}

func PrintFactorials(n int) {
	p := 1
	for i := 1; i <= n; i++ {
		fmt.Println(p)
		p *= i
	}
}

func NumProcessor() {
	fmt.Println("Parallel and concurrent programming.")

	//Go will tell us how many processors we have.
	fmt.Println("Num processors:", runtime.NumCPU())

	n := 100000000

	start := time.Now()
	Factorial(n)
	elapsed := time.Since(start)
	log.Printf("Multiprocessors took %s", elapsed)

	//we can force Go to use a different max number of procs
	runtime.GOMAXPROCS(1)

	start2 := time.Now()
	Factorial(n)
	elapsed2 := time.Since(start2)
	log.Printf("Single processor took %s", elapsed2)
}

func Factorial(n int) int {
	prod := 1
	if n == 0 {
		return 1
	}
	for i := 1; i <= n; i++ {
		prod *= i
	}
	return prod
}

//homework

func MaxMultiProcs(a []int, numProcs int) int {
	n := len(a)
	c := make(chan int)
	max := a[0]

	for i := 0; i < numProcs; i++ {
		start := i * (n / numProcs)
		end := (i + 1) * (n / numProcs)
		if i < numProcs-1 {
			go MaxOneProc(a[start:end], c)
		} else {
			go MaxOneProc(a[start:], c)
		}
	}

	for i := 0; i < numProcs; i++ {
		t := <-c
		if t > max {
			max = t
		}
	}
	return max
}

func MaxOneProc(a []int, c chan int) {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	c <- max
}

func main() {
	var wg sync.WaitGroup
	for i := range foo {
		wg.Add(1)
		go func() {
			defer wg.Done(1)
			f(bar)
		}()
	}
	wg.Wait()
}
