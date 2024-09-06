package main

import "fmt"

type Node struct {
	data any
	next *Node
}
type LinkedList struct {
	head *Node
}

func (l *LinkedList) insertAtFront(data any) {
	newNode := &Node{data: data, next: nil}
	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}

}
func (l *LinkedList) insertAtEnd(data any) {
	newNode := &Node{data: data, next: nil}

	if l.head == nil {
		l.head = newNode
		return
	}
	ptr := l.head
	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = newNode
}

func (l *LinkedList) insertAt(pos int, data any) {
	newNode := &Node{data: data, next: nil}
	ptr := l.head

	if l.head == nil {
		l.head = newNode
		return
	}

	if pos == 0 {
		newNode.next = l.head
		l.head = newNode
		return
	}
	i := 0
	for i < pos-1 {
		ptr = ptr.next
		i++
	}
	newNode.next = ptr.next
	ptr.next = newNode

}
func (l *LinkedList) deleteFirst() {
	if l.head == nil {
		fmt.Println("list is empty..")
		return
	}
	l.head = l.head.next
}

func (l *LinkedList) deleteLast() {
	ptr := l.head
	for ptr.next.next != nil {
		ptr = ptr.next
	}
	ptr.next = nil
}

func (l *LinkedList) deleteAt(pos int) {
	ptr := l.head
	i := 0
	if pos == 0 {
		l.head = l.head.next
		return
	}
	for i < pos-1 {
		ptr = ptr.next
		i++
	}
	ptr.next = ptr.next.next
}

func (l *LinkedList) deleteByValue(value any) {
	ptr := l.head

	if value == l.head.data {
		l.head = l.head.next
		return
	}
	for ptr.next.data != value {
		ptr = ptr.next
	}
	ptr.next = ptr.next.next
}

func (l *LinkedList) printList() {
	var ptr *Node = l.head
	for ptr != nil {
		fmt.Printf("%v-->", ptr.data)
		ptr = ptr.next
	}
}

func (l *LinkedList) findByIndex(pos int) any {
	ptr := l.head
	i := 0
	for i < pos {
		ptr = ptr.next
		i++
	}
	return ptr.data
}

func main() {
	myList := LinkedList{}

	myList.insertAtEnd(50)
	myList.insertAtEnd(60)
	myList.insertAtFront(40)
	myList.insertAtFront(30)
	myList.insertAt(0, 45)
	// myList.deleteByValue(45)
	fmt.Println(myList.findByIndex(4))

	myList.printList()
}
