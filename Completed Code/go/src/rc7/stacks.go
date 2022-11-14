package main

import "fmt"

type Stack []int

func (S *Stack) Push(item int){
	*S = append(*S,item)
}

func (S *Stack) Pop() int{

	//fmt.Println(S)

	if(len(*S)<1){
		panic("EMPTY STACK, CANNOT POP")
	}

	retval := (*S)[len(*S)-1]

	*S = (*S)[:len(*S)-1]

	return retval

}


func (s *Stack) ElementExists(val int) bool{
	var exists bool = false

	var temp Stack
	for true {
		if len(*s)<1{
			break
		}
		element := s.Pop()
		if element == val{
			exists = true
		}
		temp.Push(element)
	}

	for true {
		if len(temp)<1{
			break
		}
		s.Push(temp.Pop())
	}


	return exists
}



func main() {
	var s Stack
	s.Push(5)
	s.Push(7)
	//x := s.Pop()
	s.Push(11)
	//y := s.Pop()
	//fmt.Println(x)
	//fmt.Println(y)
	fmt.Println(s)
	fmt.Println(s.ElementExists(56))
}
