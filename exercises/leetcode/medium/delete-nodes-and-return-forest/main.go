package main

import "fmt"

func main() {
	r := &TreeNode{Val:1}
	r.Left = &TreeNode{Val:2}
	r.Right = &TreeNode{Val:3}
	r.Right.Right = &TreeNode{Val:4}
	forests := delNodes(r,[]int{2,1})
	for _, root := range forests {
		fmt.Println(root.Val)
	}
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
	type ToDelete struct {
		Parent *TreeNode 
		ChildPos string
	}
	toDeletes := make([]*ToDelete,0)
	for _,vtd := range to_delete {
		node, childPosition := find(root, vtd)
		td := &ToDelete{Parent: node, ChildPos: childPosition}
		toDeletes = append(toDeletes,td)
		child := &TreeNode{}
		if childPosition == "left" {
			child = node.Left
		} else if childPosition == "right" {
			child = node.Right
		} else {
			//if neither left nor right return a value then the value must be root
			child = root
		}
		if child.Left != nil {
			fr = append(fr,child.Left)
		}
		if child.Right != nil {
			fr = append(fr,child.Right)
		}
	}
	//do the actual deletion
	for _, td := range toDeletes {
		if td.ChildPos == "left" {
			td.Parent.Left = nil
		} else if td.ChildPos == "right"  {
			td.Parent.Right = nil
		}
	}
	
	if root != nil {
		fr = append(fr,root)
	}
	//remove nodes contained in to delete list
	for _, tdv := range to_delete {
		for i,n := range fr {
			if tdv == n.Val {
				fr[len(fr)-1], fr[i] = fr[i], fr[len(fr)-1]
				fr = fr[:len(fr)-1]
				break
			}
		}
	}
	return fr
}

func find(r *TreeNode, v int) (*TreeNode, string) {
	if r == nil {
		return nil, ""
	}

	if r.Left != nil {
		if r.Left.Val == v {
			return r, "left"
		} else {
			n,p := find(r.Left, v)
			if n != nil {
				return n,p
			}
		}
	}

	if r.Right != nil {
		if r.Right.Val == v {
			return r, "right"
		} else {
			n,p := find(r.Right, v)
			if n != nil {
				return n,p
			}
		}
	}
	return nil, ""
}