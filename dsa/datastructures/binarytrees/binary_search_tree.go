package binarytrees

import (
	"fmt"
)

type BinarySearchTree struct {
	root *Node
}

func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}

// Insert inserts a new node with the given data into the binary tree.
func (b *BinarySearchTree) Insert(data int) *BinarySearchTree {
	newNode := &Node{data: data}
	if b.root == nil {
		b.root = newNode
	} else {
		b.root.insert(data)
	}
	return b
}

// insert recursively finds the correct position to insert the new node.
func (n *Node) insert(data int) {
	if data <= n.data {
		if n.left == nil {
			n.left = &Node{data: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data}
		} else {
			n.right.insert(data)
		}
	}
}

// InOrder traversal (Left -> Root -> Right)
func (b *BinarySearchTree) InOrder() {
	inOrder(b.root)
	fmt.Println()
}

func inOrder(node *Node) {
	if node == nil {
		return
	}
	inOrder(node.left)
	fmt.Print(node.data, " ")
	inOrder(node.right)
}

// PreOrder traversal (Root -> Left -> Right)
func (b *BinarySearchTree) PreOrder() {
	preOrder(b.root)
	fmt.Println()
}

func preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.data, " ")
	preOrder(node.left)
	preOrder(node.right)
}

// PostOrder traversal (Left -> Right -> Root)
func (b *BinarySearchTree) PostOrder() {
	postOrder(b.root)
	fmt.Println()
}

func postOrder(node *Node) {
	if node == nil {
		return
	}
	postOrder(node.left)
	postOrder(node.right)
	fmt.Print(node.data, " ")
}

// Print prints the binary tree in a structured way.
func (b *BinarySearchTree) Print(node *Node, space int, child rune) {
	if node == nil {
		return
	}

	// Print the current node with indentation
	fmt.Printf("%s%c:%d\n", fmt.Sprintf("%*s", space, ""), child, node.data)

	// Recursively print left and right subtrees with additional indentation
	b.Print(node.left, space+4, 'L')
	b.Print(node.right, space+4, 'R')
}

func BST() {
	tree := NewBinarySearchTree()
	tree.Insert(8).
		Insert(1).
		Insert(3).
		Insert(6).
		Insert(9).
		Insert(10).
		Insert(4).
		Insert(5).
		Insert(8).
		Insert(10).
		Print(tree.root, 0, 'M')
}
