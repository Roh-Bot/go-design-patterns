package linkedlists

type CircularSinglyNode struct {
	data int
	next *CircularSinglyNode
}

type CircularSinglyList struct {
	head *CircularSinglyNode
	tail *CircularSinglyNode
}

func NewCircularSinglyList() *CircularSinglyList {
	return new(CircularSinglyList)
}

func (c *CircularSinglyList) Append(data int) {
	newNode := &CircularSinglyNode{data: data, next: nil}
	if c.head == nil {
		newNode.next = newNode
		c.head = newNode
		c.tail = newNode
		return
	}

	newNode.next = c.head
	c.tail.next = newNode
	c.tail = newNode
}

func (c *CircularSinglyList) Prepend(data int) {
	newNode := &CircularSinglyNode{data: data, next: c.head}
	c.head = newNode
	c.tail.next = newNode
}

func (c *CircularSinglyList) Delete(data int) {
	if c.head == nil {
		return
	}

	if c.head.data == data {
		if c.head == c.tail {
			c.head = nil
			c.tail = nil
		} else {
			c.head = c.head.next
			c.tail.next = c.head
		}
		return
	}

	current := c.head
	for {
		if current.next.data == data {
			if current.next == c.tail {
				c.tail = current
				c.tail.next = c.head
			} else {
				current.next = current.next.next
			}
			return
		}
		current = current.next
		if current == c.head {
			break
		}
	}
}

func (c *CircularSinglyList) Reverse() {
	current := c.head
	var nextNode *CircularSinglyNode
	var prevNode *CircularSinglyNode
	c.tail = current
	for {
		nextNode = current.next
		current.next = prevNode
		prevNode = current
		current = nextNode
		if current == c.head {
			c.tail.next = prevNode
			break
		}
	}
	c.head = prevNode
}

func CSinglyLinkedList() {
	c := NewCircularSinglyList()
	c.Append(10)
	c.Append(20)
	c.Append(30)
	c.Append(40)
	c.Reverse()
}
