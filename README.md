# Golang Algo
This is a repo mainly focus on study, learning and practices golang with algo

## Binary Search Tree - bst.go
Create and remove the node from the node, print out the value of the node from small to large 

### How to use
#### Run
 - go run bst.go

## Find the Distance between Two Nodes of a Binary Tree - btd.go
This is a golang program to find the distance between 2 nodes of a binary tree.

                   [5]
               /          \
            [10]           [15]
            /    \         /    \
       [20]      [25]  [30]     [35]
       /   \     /  \      \    
     [40] [45] [50] [55]   [60]

### Formula
 Distance(x,y) = (Distance(root,x) + Distance(root,y)) - 2 * Distance(root, Lowest Common Ancestor(x,y) )

### How to use
#### Compile
 - go build btd.go
#### Run
 - ./btd 20 35

### Sample Result
----------------------
Node 1:  45
Node 2:  55
Distance:  4
----------------------
Node 1:  45
Node 2:  60
Distance:  6