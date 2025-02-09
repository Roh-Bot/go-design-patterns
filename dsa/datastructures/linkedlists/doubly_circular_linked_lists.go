package linkedlists

type CircularDoublyNode struct {
	prev *CircularDoublyNode
	data int
	next *CircularDoublyNode
}

type CircularDoubleLinkedList struct {
	head *CircularDoublyNode
	tail *CircularDoublyNode
}

func NewCircularDoublyLinkedList() *CircularDoubleLinkedList {
	return new(CircularDoubleLinkedList)
}

func (c *CircularDoubleLinkedList) Append(data int) {
	newNode := &CircularDoublyNode{data: data}
	if c.head == nil {
		newNode.next = newNode
		newNode.prev = newNode
		c.head = newNode
		c.tail = newNode
		return
	}

	newNode.next = c.head
	newNode.prev = c.tail
	c.tail.next = newNode
	c.head.prev = newNode
	c.tail = newNode
}

func (c *CircularDoubleLinkedList) Prepend(data int) {
	newNode := &CircularDoublyNode{data: data}
	newNode.next = c.head
	newNode.prev = c.tail
	c.head = newNode
}

func (c *CircularDoubleLinkedList) Delete(data int) {
	if c.head == c.tail && c.head.data == data {
		c.head = nil
		c.tail = nil
		return
	}

	if c.head.data == data {
		c.head.next.prev = c.tail
		c.head = c.head.next
		c.tail.next = c.head
		return
	}

	if c.tail.data == data {
		c.tail.prev.next = c.head
		c.tail = c.tail.prev
		c.head.prev = c.tail
		return
	}

	current := c.head
	for current.next != c.tail {
		if current.next.data == data {
			current.next.next.prev = current
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func (c *CircularDoubleLinkedList) Reverse() {
	current := c.head
	var prev *CircularDoublyNode
	var next *CircularDoublyNode
	c.tail = current
	for {
		next = current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
		if current == c.head {
			c.tail.next = prev
			break
		}
	}
	c.head = prev
}
func DDoublyLinkedList() {
	d := NewCircularDoublyLinkedList()
	d.Append(10)
	d.Append(20)
	d.Append(30)
	d.Append(40)
	d.Delete(10)
}
