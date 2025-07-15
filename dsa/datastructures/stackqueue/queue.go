package stackqueue

import "fmt"

// Queue represents a generic queue.
type Queue[T comparable] []T

// IsEmpty checks if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

// Enqueue adds a new element to the end of the queue.
func (q *Queue[T]) Enqueue(data T) {
	*q = append(*q, data) // Append the new element to the end of the queue
}

// Dequeue removes and returns the element from the front of the queue. Returns false if the queue is empty.
func (q *Queue[T]) Dequeue() (T, bool) {
	var zeroVal T
	if q.IsEmpty() {
		return zeroVal, false
	}
	e := (*q)[0]
	*q = (*q)[1:] // Remove the front element from the queue by slicing
	return e, true
}

// Peek returns the front element of the queue without removing it. Returns false if the queue is empty.
func (q *Queue[T]) Peek() (T, bool) {
	var zeroValue T
	if q.IsEmpty() {
		return zeroValue, false
	}
	return (*q)[0], true
}

// Length returns the number of elements in the queue.
func (q *Queue[T]) Length() int {
	return len(*q)
}

// Clear removes all elements from the queue.
func (q *Queue[T]) Clear() {
	*q = nil
}

func Queue1() {
	queue := new(Queue[string]) // Create a queue variable of type Queue

	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")

	fmt.Println("Queue contents:")

	queue.Enqueue("new")
	queue.Enqueue("elements")

	e, _ := queue.Peek()
	fmt.Println("Peek at front element: ", e)
	fmt.Println("Queue size:", queue.Length())

	queue.Clear()
	fmt.Println("Queue cleared. Is empty?", queue.IsEmpty())
}
