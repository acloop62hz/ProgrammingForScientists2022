package main

import "fmt"

func main() {
	c := make(chan string)
	c <- "hello"
	fmt.Println(<-c)
}
