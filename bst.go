/*
Binary Search Tree
*/

package main

import "fmt"

type Node struct {
	value	int
	left	*Node
	right 	*Node
}

var node *Node

func main() {
	insert(50)
	insert(30)
	insert(20)
	insert(70)
	insert(80)
	insert(40)
	delete(50)
	print()
}

func insert(value int) {
	node = insertValue(node, value)
}

func delete(value int) {
	node = deleteValue(node, value)
}

func print() {
	printNode(node)
}

func insertValue(node *Node, value int) *Node {
	if node == nil {
		return &Node{value, nil, nil}
	}

	if value < node.value {
		node.left = insertValue(node.left, value)
	} else {
		node.right = insertValue(node.right, value)
	}

	return node
}

func deleteValue(node *Node, value int) *Node {
	if node == nil {
		return node
	}

	if value < node.value {
		node.left = deleteValue(node.left, value)
	} else if value > node.value {
		node.right = deleteValue(node.right, value)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			node.value = findMinValue(node.right)
			node.right = deleteValue(node.right, node.value)
		}
	}
	return node
}

func findMinValue(node *Node) int {
	if node == nil {
		return 0
	} else if node.left == nil {
		return node.value
	} else {
		return findMinValue(node.left)
	}
}

func printNode(node *Node) {
	if node != nil {
		printNode(node.left)
		fmt.Println("value = ", node.value)
		printNode(node.right)
	}
}