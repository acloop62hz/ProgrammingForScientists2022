// Linked List in Golang
package main

import "fmt"

type Node struct {
	next *Node
	key  int
}

type List struct {
	head *Node
}

func (L *List) InsertBeginning(key int) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	L.head = list
}

func (L *List) DeleteBeginning() {
	L.head = L.head.next
}



func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}



func main() {
	link := List{}
	link.InsertBeginning(5)
	link.InsertBeginning(9)
	link.InsertBeginning(13)
	link.InsertBeginning(22)
	link.InsertBeginning(28)
	link.InsertBeginning(36)
	fmt.Println(link,link.head, link.head.next)
	fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", link.head.key)
	link.Display()
	link.DeleteBeginning()
	fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", link.head.key)
	link.Display()
	fmt.Println("\n==============================\n")
}
