package main

import (
	"fmt"
)

func main() {
	linkedList := doublyLinkedList{}
	linkedList.insertAtFront(10)
	linkedList.insertAtFront(5)
	linkedList.insertAtEnd(15)
	linkedList.insertAt(5, 7)

	fmt.Println("List In Forward Manner.")
	linkedList.printForward()
	fmt.Println()
	fmt.Println("List in Backword Manner.")
	linkedList.printBackward()
	fmt.Println()
	linkedList.deleteitem(10)
	fmt.Println("List In Forward Manner after deletion .")
	linkedList.printForward()
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

func (d *doublyLinkedList) insertAt(pos int, data int) {
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
	if pos == 0 {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
		return
	}

	ptr := d.head
	n := 1
	for n < pos && ptr.next != nil {
		ptr = ptr.next
		n++
	}

	if ptr.next == nil {
		ptr.next = newNode
		newNode.prev = ptr
		d.tail = newNode
		return
	}

	newNode.next = ptr.next
	newNode.prev = ptr
	if ptr.next != nil {
		ptr.next.prev = newNode
	}
	ptr.next = newNode
}

func (d *doublyLinkedList) deleteitem(data int) {
	ptr := d.head

	if d.head.data == data {
		d.head = d.head.next
		return
	}
	if d.tail.data == data {
		d.tail = d.tail.prev
		d.tail.next = nil
		return
	}
	for ptr.data != data {
		ptr = ptr.next
	}
	ptr.prev.next = ptr.next
	ptr.next.prev = ptr.prev
	ptr.next = nil
	ptr.prev = nil
}
func (d *doublyLinkedList) printForward() {
	ptr := d.head
	for ptr != nil {
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
