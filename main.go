package main

import "fmt"

func main() {
	arr := []string{"duh", "ill"}
	for _, str := range arr {
		var sum int32
		for _, v := range str {
			sum += v
		}
		fmt.Println(sum)

	}
}

type node struct {
	left  *node
	data  int
	right *node
}

type BinaryTree struct {
	root *node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (b *BinaryTree) PreOrder() {
	fmt.Println(b.root)
}
