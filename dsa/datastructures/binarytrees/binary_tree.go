package binarytrees

import (
	"fmt"

	"github.com/RohBot/GOLANG/GoPhase3/dsa/datastructures/stackqueue"
)

// BinaryTree defines a binary tree with a pointer to its root node.
type BinaryTree struct {
	Root *Node
}

// NewBinaryTree creates and returns a new empty binary tree.
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

// PreOrder performs a recursive preorder traversal (Root → Left → Right).
func (b *BinaryTree) PreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Data) // Visit root
	b.PreOrder(node.Left)  // Traverse left subtree
	b.PreOrder(node.Right) // Traverse right subtree
}

// PreOrderIterative performs an iterative preorder traversal using a stack.
// It visits nodes in Root → Left → Right order.
func (b *BinaryTree) PreOrderIterative() []int {
	stack := stackqueue.Stack[*Node]{}
	stack.Push(b.Root)
	var result []int

	for !stack.IsEmpty() {
		v := stack.Pop()
		result = append(result, v.Data) // Visit root

		// Push right child first so left is processed first
		if v.Right != nil {
			stack.Push(v.Right)
		}
		if v.Left != nil {
			stack.Push(v.Left)
		}
	}
	return result
}

// InOrderIterative performs an iterative inorder traversal (Left → Root → Right).
// It uses a stack to mimic the recursive call stack.
func (b *BinaryTree) InOrderIterative() []int {
	stack := stackqueue.Stack[*Node]{}
	node := b.Root
	var result []int

	// Loop to go left until it reaches the leaf node
	// Once leaf node is found check if the stack is not empty
	// for loop termination and pop the stack to append the value in the result
	// After the pop simply assign the node to the right node of the
	// corresponding node
	for {
		if node != nil {
			stack.Push(node) // Go left as far as possible
			node = node.Left
			continue
		}
		if stack.IsEmpty() {
			break
		}
		val := stack.Pop()
		result = append(result, val.Data) // Visit node
		node = val.Right                  // Then go right
	}
	return result
}

// PostOrderIterative performs an iterative postorder traversal (Left → Right → Root).
// This version does reverse of root → right → left and reverses the result.
func (b *BinaryTree) PostOrderIterative() []int {
	stack := stackqueue.Stack[*Node]{}
	node := b.Root
	var result []int

	// This is a variant that visits root → right → left,
	// and we'll reverse the result at the end to get postorder
	for {
		if node != nil {
			stack.Push(node)
			result = append([]int{node.Data}, result...) // Prepend (reverse order)
			node = node.Right                            // Go right first
			continue
		}
		if stack.IsEmpty() {
			break
		}
		val := stack.Pop()
		node = val.Left // After right, go left
	}
	return result
}

// LevelOrderIterative performs a standard level-order traversal (Breadth-First Search).
// It uses a queue to process nodes level by level.
func (b *BinaryTree) LevelOrderIterative() {
	queue := stackqueue.Queue[*Node]{}
	queue.Enqueue(b.Root)
	var result []int

	for !queue.IsEmpty() {
		node, _ := queue.Dequeue()
		result = append(result, node.Data) // Visit node

		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}

	fmt.Println(result)
}

// LevelOrderGrouped performs level-order traversal but groups nodes by level.
// It returns a 2D slice: each sub-slice contains node values at one level.
func (b *BinaryTree) LevelOrderGrouped() [][]int {
	if b.Root == nil {
		return nil
	}

	queue := stackqueue.Queue[*Node]{}
	queue.Enqueue(b.Root)
	var result [][]int

	for !queue.IsEmpty() {
		levelSize := queue.Length() // Get the size of the queue which actually infers to the size of the current level
		level := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node, _ := queue.Dequeue()
			level = append(level, node.Data)

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
		result = append(result, level)
	}

	return result
}

func (b *BinaryTree) MaximumDepth(node *Node) int {
	if b.Root == nil {
		return 0
	}

	lh := b.MaximumDepth(b.Root.Left)
	rh := b.MaximumDepth(b.Root.Right)

	return 1 + max(lh, rh)
}
