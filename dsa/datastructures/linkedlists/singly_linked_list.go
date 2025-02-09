package linkedlists

import (
	"errors"
	"fmt"
)

// SinglyNode represents a single node in the linked list.
type SinglyNode struct {
	data int
	next *SinglyNode
}

// SinglyLinkedList represents the linked list.
type SinglyLinkedList struct {
	head *SinglyNode
	tail *SinglyNode
}

// NewLinkedList creates and returns a new empty linked list.
func NewLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// Append adds a new node with the given data to the end of the list.
func (s *SinglyLinkedList) Append(data int) {
	newNode := &SinglyNode{data: data, next: nil}

	if s.head == nil {
		s.head = newNode
		s.tail = newNode
		return
	}
	s.tail.next = newNode
	s.tail = newNode
}

// Prepend adds a new node with the given data to the beginning of the list.
func (s *SinglyLinkedList) Prepend(data int) {
	newNode := &SinglyNode{data: data, next: s.head}
	s.head = newNode
}

// Delete removes the first occurrence of a node with the given data from the list.
func (s *SinglyLinkedList) Delete(data int) {
	if s.head == nil {
		return
	}

	if s.head.data == data {
		if s.head == s.tail {
			s.head = nil
			s.tail = nil
			return
		}
		s.head = s.head.next
		return
	}

	current := s.head
	for current.next != nil {
		switch {
		case current.next.data == data:
			current.next = current.next.next
			return
		case current.next == s.tail && current.next.data == s.tail.data:
			current.next = nil
			s.tail = current
			return
		}
		current = current.next
	}
}

// Find searches for a node with the given data and returns true if found.
func (s *SinglyLinkedList) Find(data int) bool {
	current := s.head
	for current != nil {
		if current.data == data {
			return true
		}
		current = current.next
	}
	return false
}

// Length returns the number of nodes in the list.
func (s *SinglyLinkedList) Length() int {
	count := 0
	current := s.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

// Reverse reverses the order of nodes in the list.
func (s *SinglyLinkedList) Reverse() {
	current := s.head
	var prev *SinglyNode
	var next *SinglyNode
	s.tail = current
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	s.head = prev
}

func (s *SinglyLinkedList) GetNthNodeFromEnd(n int) *SinglyNode {
	first := s.head
	second := s.head

	for range n {
		first = first.next
	}

	for first != nil {
		first = first.next
		second = second.next
	}
	return second
}

func (s *SinglyLinkedList) GetNodeAtIndex(index int) (*SinglyNode, error) {
	current := s.head
	for i := 0; current != nil; i++ {
		if index == i {
			return current, nil
		}
		current = current.next
		index++
	}
	return nil, errors.New(fmt.Sprintf(" index out of range [%d]", index))
}

func (s *SinglyLinkedList) LoopDetector() bool {
	fast := s.head
	slow := s.head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if slow == fast {
			return true
		}
	}
	return false
}

// Display prints the linked list in a readable format.
func (s *SinglyLinkedList) Display() {
	current := s.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (s *SinglyLinkedList) CreateLoop(data int) {
	newNode := &SinglyNode{data: data, next: s.head.next}

	current := s.head

	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func SLinkedList() {
	s := NewLinkedList()

	s.Append(10)
	s.Append(20)
	s.Append(30)
	s.Append(40)
	s.Append(50)
	s.Reverse()
	s.GetNodeAtIndex(5)
	//s.CreateLoop(60)
	//fmt.Println(s.LoopDetector())
	//fmt.Println(s.GetNthNodeFromEnd(2))

	fmt.Println("Linked List:")
	s.Display()

	//s.Prepend(5)
	//fmt.Println("After Prepend(5):")
	//s.Display()

	s.Delete(30)
	fmt.Println("After Delete(20):")
	s.Display()

	fmt.Printf("Find 10: %v\n", s.Find(10))
	fmt.Printf("Find 20: %v\n", s.Find(20))

	fmt.Printf("Length: %d\n", s.Length())

	s.Reverse()
	fmt.Println("After Reverse:")
	s.Display()
}
