package main

import "fmt"

func main() {
	r := &TreeNode{Val:1}

	fmt.Println(delNodes(r,[]int{3,5}))
}

/*
	* Given the root of a binary tree, each node in the tree has a distinct value.

	* After deleting all nodes with a value in to_delete, we are left with a forest (a disjoint union of trees).

	* Return the roots of the trees in the remaining forest.  You may return the result in any order.
*/
type TreeNode struct {
	Val int 
	Left *TreeNode 
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	fr := make([]*TreeNode,0)
	for _,vtd := range to_delete {
			fmt.Println(vtd)
	}
	
	return fr
}