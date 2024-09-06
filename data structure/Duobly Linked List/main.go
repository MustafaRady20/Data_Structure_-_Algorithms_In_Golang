package main

import (
	"fmt"
)

func main() {
	linkedList := doublyLinkedList{}
	linkedList.insertAtFront(10)
	linkedList.insertAtFront(5)
	linkedList.insertAtEnd(15)
	fmt.Println("List In Forward Manner.")
	linkedList.printForward()
	fmt.Println()
	fmt.Println("List in Backword Manner.")
	linkedList.printBackward()
	fmt.Println()

}

type Node struct {
	data int
	next *Node
	prev *Node
}

type doublyLinkedList struct {
	head *Node
	tail *Node
}

func (d *doublyLinkedList) insertAtFront(data int) {
	newNode := &Node{
		data: data,
		next: nil,
		prev: nil,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}
	newNode.next = d.head
	d.head.prev = newNode
	d.head = newNode

}

func (d *doublyLinkedList) insertAtEnd(data int) {
	newNode := &Node{
		data: data,
		next: nil,
		prev: nil,
	}
	if d.tail == nil {
		d.head = newNode
		d.tail = newNode
		return
	}

	d.tail.next = newNode
	newNode.prev = d.tail
	d.tail = newNode

}
func (d *doublyLinkedList) printForward() {
	ptr := d.head
	for ptr != nil {
		// if ptr.prev !=nil{
		// 	fmt.Printf("<-(%d)->", ptr.prev.data)
		// }
		fmt.Printf("%d-->", ptr.data)
		ptr = ptr.next
	}
}

func (d *doublyLinkedList) printBackward() {
	ptr := d.tail
	if d.tail == nil {
		fmt.Println("List is empty")
		return
	}
	for ptr != nil {
		fmt.Printf("%d-->", ptr.data)
		ptr = ptr.prev
	}

}
