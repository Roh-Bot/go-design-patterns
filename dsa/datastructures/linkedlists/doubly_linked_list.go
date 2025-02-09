package linkedlists

import (
	"fmt"
)

// DoublyNode represents a single node in the doubly linked list.
type DoublyNode struct {
	prev *DoublyNode
	data int
	next *DoublyNode
}

// DoublyLinkedList represents the doubly linked list.
type DoublyLinkedList struct {
	head *DoublyNode
	tail *DoublyNode
}

// NewDoublyLinkedList creates and returns a new empty doubly linked list.
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Append adds a new node with the given data to the end of the list.
func (d *DoublyLinkedList) Append(data int) {
	newNode := &DoublyNode{data: data, prev: nil, next: nil}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}

	newNode.prev = d.tail
	d.tail.next = newNode
	d.tail = newNode
}

// Prepend adds a new node with the given data to the beginning of the list.
func (d *DoublyLinkedList) Prepend(data int) {
	newNode := &DoublyNode{data: data, prev: nil, next: d.head}

	if d.head != nil {
		d.head.prev = newNode
	}
	d.head = newNode

	if d.tail == nil {
		d.tail = newNode
	}
}

// Delete removes the first occurrence of a node with the given data from the list.
func (d *DoublyLinkedList) Delete(data int) {
	if d.head.data == data {
		d.head = d.head.next
		d.head.prev = nil
	}

	if d.tail.data == data {
		d.tail = d.tail.prev
		d.tail.next = nil
	}

	current := d.head.next
	for current != nil {
		if current.data == data {
			current.next.prev = current.prev
			current.prev.next = current.next
			return
		}
		current = current.next
	}
}

// Find searches for a node with the given data and returns true if found.
func (d *DoublyLinkedList) Find(data int) bool {
	current := d.head
	for current != nil {
		if current.data == data {
			return true
		}
		current = current.next
	}
	return false
}

// Length returns the number of nodes in the list.
func (d *DoublyLinkedList) Length() int {
	count := 0
	current := d.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

// Reverse reverses the order of nodes in the list.
func (d *DoublyLinkedList) Reverse() {
	var prev *DoublyNode
	current := d.head
	d.tail = current

	for current != nil {
		next := current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}
	d.head = prev
}

// DisplayForward prints the doubly linked list from head to tail.
func (d *DoublyLinkedList) DisplayForward() {
	current := d.head
	for current != nil {
		fmt.Printf("%d <-> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

// DisplayBackward prints the doubly linked list from tail to head.
func (d *DoublyLinkedList) DisplayBackward() {
	current := d.tail
	for current != nil {
		fmt.Printf("%d <-> ", current.data)
		current = current.prev
	}
	fmt.Println("nil")
}

func DLinkedList() {
	d := NewDoublyLinkedList()

	d.Append(10)
	d.Append(20)
	d.Append(30)
	d.Append(40)

	d.Reverse()

	fmt.Println("Doubly Linked List (Forward):")
	d.DisplayForward()

	fmt.Println("Doubly Linked List (Backward):")
	d.DisplayBackward()

	d.Prepend(5)
	fmt.Println("After Prepend(5):")
	d.DisplayForward()

	d.Delete(20)
	fmt.Println("After Delete(20):")
	d.DisplayForward()

	fmt.Printf("Find 10: %v\n", d.Find(10))
	fmt.Printf("Find 20: %v\n", d.Find(20))

	fmt.Printf("Length: %d\n", d.Length())

	d.Reverse()
	fmt.Println("After Reverse:")
	d.DisplayForward()
}
