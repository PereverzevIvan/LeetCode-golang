package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodes := map[int]*TreeNode{}
	hasParent := map[int]bool{}

	for _, desc := range descriptions {
		parent := desc[0]
		child := desc[1]
		isLeft := desc[2]

		if _, inMap := nodes[parent]; !inMap {
			nodes[parent] = &TreeNode{parent, nil, nil}
		}

		if _, inMap := nodes[child]; !inMap {
			nodes[child] = &TreeNode{child, nil, nil}
		}

		if isLeft == 1 {
			nodes[parent].Left = nodes[child]
		} else {
			nodes[parent].Right = nodes[child]
		}

		hasParent[child] = true
	}

	root := 0
	for _, n := range nodes {
		// fmt.Printf("Узел %v имеет детей %v и %v\n", n.Val, n.Left, n.Right)
		if _, inMap := hasParent[n.Val]; !inMap {
			root = n.Val
		}
	}

	return nodes[root]
}

func main() {
	desc := [][]int{
		[]int{39, 70, 1},
		[]int{13, 39, 1},
		[]int{85, 74, 1},
		[]int{74, 13, 1},
		[]int{38, 82, 1},
		[]int{82, 85, 1},
	}
	fmt.Println(createBinaryTree(desc))
}
