package main

import "fmt"

func main() {
	r := &TreeNode{Val:1}
	r.Left = &TreeNode{Val:2}
	r.Right = &TreeNode{Val:3}
	r.Right.Right = &TreeNode{Val:4}
	forests := delNodesOptimised(r,[]int{2,1})
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

/*
 * A more efficient approach [https://leetcode.com/problems/delete-nodes-and-return-forest/discuss/357357/Python-O(n)-DFS-Recursion-with-Explanation]
	
 	* Make the to_delete a hashset for O(1) look-up, check if the root needs to be included in the result or not, and proceed with DFS. For DFS we need to check if the specified node needs to be deleted, if yes, we need to update the parent node pointers and add in it's children to the result if they are not in the delete list. Then continue DFS keeping track of parent and where you came from (R or L).
*/
func delNodesOptimised(root *TreeNode, to_delete []int) []*TreeNode {
	forests := make([]*TreeNode,0)
	toDelete := make(map[int]bool)
	for _,v := range to_delete {
		toDelete[v] = true
	}
	forests = append(forests, delNodesOptimisedHelper(root,&toDelete)...)
	if _,ok := toDelete[root.Val]; !ok {
		forests = append(forests,root)
	}
	return forests
}

func delNodesOptimisedHelper(root *TreeNode, toDelete *map[int]bool) []*TreeNode {
	forests := make([]*TreeNode,0)
	//deal with children first
	if root.Left != nil {
		forests = append(forests, delNodesOptimisedHelper(root.Left,toDelete)...)
	} 

	if root.Right != nil {
		forests = append(forests, delNodesOptimisedHelper(root.Right,toDelete)...)
	} 

	if _,ok := (*toDelete)[root.Val]; ok {
		//this value is to be deleted
		if root.Left != nil {
			if _,ok := (*toDelete)[root.Left.Val]; !ok {
				forests = append(forests,root.Left)
			} else {
				root.Left = nil
			}
		} 

		if root.Right != nil {
			if _,ok := (*toDelete)[root.Right.Val]; !ok {
				forests = append(forests,root.Right)
			} else {
				root.Right = nil
			}
		}
	} else {
		if root.Left != nil {
			if _,ok := (*toDelete)[root.Left.Val]; ok {
				root.Left = nil
			}
		} 
		if root.Right != nil {
			if _,ok := (*toDelete)[root.Right.Val]; ok {
				root.Right = nil
			}
		} 
	}
	return forests
}