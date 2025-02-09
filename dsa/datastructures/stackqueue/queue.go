package stackqueue

import "fmt"

// Queue represents a queue of strings.
type Queue []string

// IsEmpty checks if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Enqueue adds a new element to the end of the queue.
func (q *Queue) Enqueue(str string) {
	*q = append(*q, str) // Append the new element to the end of the queue
}

// Dequeue removes and returns the element from the front of the queue. Returns false if the queue is empty.
func (q *Queue) Dequeue() bool {
	if q.IsEmpty() {
		return false
	}
	*q = (*q)[1:] // Remove the front element from the queue by slicing
	return true
}

// Peek returns the front element of the queue without removing it. Returns false if the queue is empty.
func (q *Queue) Peek() (string, bool) {
	if q.IsEmpty() {
		return "", false
	}
	return (*q)[0], true
}

// Size returns the number of elements in the queue.
func (q *Queue) Size() int {
	return len(*q)
}

// Clear removes all elements from the queue.
func (q *Queue) Clear() {
	*q = nil
}

func Queue1() {
	queue := new(Queue) // Create a queue variable of type Queue

	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")

	fmt.Println("Queue contents:")

	queue.Enqueue("new")
	queue.Enqueue("elements")

	e, _ := queue.Peek()
	fmt.Println("Peek at front element: ", e)
	fmt.Println("Queue size:", queue.Size())

	queue.Clear()
	fmt.Println("Queue cleared. Is empty?", queue.IsEmpty())
}
