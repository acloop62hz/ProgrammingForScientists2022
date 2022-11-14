package main

import "fmt"

type Queue []int

func (Q *Queue) Enqueue(item int) {
	*Q = append(*Q, item)
}

func (Q *Queue) Dequeue() int {

	//fmt.Println(S)

	if len(*Q) < 1 {
		panic("EMPTY QUEUE, CANNOT DEQUEUE")
	}

	retval := (*Q)[0]

	*Q = (*Q)[1:len(*Q)]

	return retval

}

func (q *Queue) ElementExists(val int) bool {
	var exists bool = false

	var temp Queue
	for true {
		if len(*q) < 1 {
			break
		}
		element := q.Dequeue()
		if element == val {
			exists = true
		}
		temp.Enqueue(element)
	}

	for true {
		if len(temp) < 1 {
			break
		}
		q.Enqueue(temp.Dequeue())
	}

	return exists
}

func main() {
	var q Queue
	q.Enqueue(5)
	q.Enqueue(7)
	x := q.Dequeue()
	q.Enqueue(11)
	y := q.Dequeue()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(q)

}
