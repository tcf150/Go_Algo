/*
Find the Distance between Two Nodes of a Binary Tree
This is a go program to find the distance between 2 nodes of a binary tree.

               [5]
           /          \
        [10]           [15]
        /    \         /    \
   [20]      [25]  [30]     [35]
   /   \     /  \      \    
 [40] [45] [50] [55]   [60]

Formula
Distance(x,y) = (Distance(root,x) + Distance(root,y)) - 2 * Distance(root, Lowest Common Ancestor(x,y) )

How to use
go build btd.go
./btd 20 35

Sample Result
----------------------
Node 1:  45
Node 2:  55
Distance:  4
----------------------
Node 1:  45
Node 2:  60
Distance:  6
*/

package main 

import "fmt"
import "os"
import "strconv"

type Node struct {
	value	int
	left	*Node
	right 	*Node
}

func main() {
	var node40 = &Node{40, nil, nil}
	var node45 = &Node{45, nil, nil}
	var node20 = &Node{20, node40, node45}
	var node50 = &Node{50, nil, nil}
	var node55 = &Node{55, nil, nil}
	var node25 = &Node{25, node50, node55}
	var node10 = &Node{10, node20, node25}
	var node60 = &Node{60, nil, nil}
	var node30 = &Node{30, nil, node60}
	var node35 = &Node{35, nil, nil}
	var node15 = &Node{15, node30, node35}
	var root = &Node{5, node10, node15}

	var node1 = convertStringToInt(os.Args[1])
	var node2 = convertStringToInt(os.Args[2])
	fmt.Println("Node 1: ", node1)
	fmt.Println("Node 2: ", node2)
	var distance = findDistanceBetweenNode(root, node1, node2)
	fmt.Println("Distance: ", distance)
}

func convertStringToInt(str string) int {
	if i, err := strconv.Atoi(str); err == nil {
		return i
	}
	return 0
}

func findDistanceBetweenNode(root *Node, valueX int, valueY int) int {
	var x = getPathLength(root, valueX) - 1
	var y = getPathLength(root, valueY) - 1
	var lca = findLowestCommonAncestor(root, valueX, valueY)
	var distanceLca = 0
	if lca != nil {
		distanceLca = getPathLength(root, lca.value) - 1
	}

	return (x + y) - 2 * distanceLca
}

func findLowestCommonAncestor(root *Node, valueX int, valueY int) *Node {
	if root == nil {
		return nil
	}
	if root.value == valueX || root.value == valueY {
		return root
	}

	var left = findLowestCommonAncestor(root.left, valueX, valueY)
	var right = findLowestCommonAncestor(root.right, valueX, valueY)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	} else {
		return nil
	}
}

func getPathLength(root *Node, value int) int {
	var x = 0
	if root.value == value {
		return x + 1
	} else {
		x = 0
		if root.left != nil {
			x = getPathLength(root.left, value)
		}

		if x > 0 {
			return x + 1
		}

		x = 0
		if root.right != nil {
			x = getPathLength(root.right, value)
		}

		if x > 0 {
			return x + 1
		}
	}
	return 0
}