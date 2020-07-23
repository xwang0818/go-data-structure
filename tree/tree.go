/* Copyright (C) Xiang Wang - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Xiang Wang <xwang1314@gmail.com>, July 2020
 */

package tree

import "container/list"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func preOrderRecursive(node *TreeNode, arr []int) []int {
    if node == nil {
        return arr
    }
    arr = append(arr, node.Val)
    arr = preOrderRecursive(node.Left, arr)
    arr = preOrderRecursive(node.Right, arr)
    return arr
}

// Use one stack
// push root
// while not empty, pop and push R and push L
func preOrderIterative(node *TreeNode, arr []int) []int {
    stack := list.New()
    stack.PushFront(node)
    for ; stack.Len() != 0; {
        front := stack.Front()
        curr := stack.Remove(front).(*TreeNode)
        arr = append(arr, curr.Val)
        if curr.Right != nil {
            stack.PushFront(curr.Right)
        }
        if curr.Left != nil {
            stack.PushFront(curr.Left)
        }
    }
    return arr
}

func inOrderRecursive(node *TreeNode, arr []int) []int {
    if node == nil {
        return arr
    }
    arr = inOrderRecursive(node.Left, arr)
    arr = append(arr, node.Val)
    arr = inOrderRecursive(node.Right, arr)
    return arr
}

// Use one stack, set curr = root
// while not empty or curr not nil
//      while curr not nil
//             go to the most Left node
//      Now curr is nil, pop from stack
//      Visit the node
//      set curr = Right node and repeat
func inOrderIterative(node *TreeNode, arr []int) []int {
    stack := list.New()
    curr := node
    for ; curr != nil || stack.Len() != 0; {
        // move curr all the way to the most Left node
        for ; curr != nil; {
            stack.PushFront(curr)
            curr = curr.Left
        }
        // curr is nil so pop from stack
        front := stack.Front()
        curr = stack.Remove(front).(*TreeNode)
        // visit the node
        arr = append(arr, curr.Val)
        // move on to right node
        curr = curr.Right
    }
    return arr
}


func postOrderRecursive(node *TreeNode, arr []int) []int {
    if node == nil {
        return arr
    }
    arr = postOrderRecursive(node.Left, arr)
    arr = postOrderRecursive(node.Right, arr)
    arr = append(arr, node.Val)
    return arr
}

// Use two stacks
// Push root into stack_one
// while stack_one not empty
//       curr = pop from stack_one
//       push curr into stack_two
//       if curr has Left, push Left into stack_one
//       if curr has Right, push Right into stack_one
// stack_two will be in postOrder
func postOrderIterative(node *TreeNode, arr []int) []int {
    stack_one := list.New()
    stack_two := list.New()
    stack_one.PushFront(node)
    for ; stack_one.Len() != 0; {
        front := stack_one.Front()
        curr := stack_one.Remove(front).(*TreeNode)
        stack_two.PushFront(curr)
        if curr.Left != nil {
            stack_one.PushFront(curr.Left)
        }
        if curr.Right != nil {
            stack_one.PushFront(curr.Right)
        }
    }
    for ; stack_two.Len() != 0; {
        front := stack_two.Front()
        curr := stack_two.Remove(front).(*TreeNode)
        arr = append(arr, curr.Val)
    }
    return arr
}


/*
* ======================== levelOrder 2D Array ========================
*/
func levelOrderRecusive(root *TreeNode) [][]int {
    var result [][]int
    if root == nil {
        return result
    }
    queue := list.New()
    queue.PushBack(root)
    return levelOrderRecusiveHelper(queue, result)
}

func levelOrderRecusiveHelper(queue *list.List, result [][]int) [][]int {
    if queue.Len() == 0 {
        return result
    }
    nextq := list.New()
    var arr []int
    for ; queue.Len() !=0; {
        front := queue.Front()
        curr := queue.Remove(front).(*TreeNode)
        arr = append(arr, curr.Val)

        if curr.Left != nil {
            nextq.PushBack(curr.Left)
        }
        if curr.Right != nil {
            nextq.PushBack(curr.Right)
        }
    }
    result = append(result, arr)
    return levelOrderRecusiveHelper(nextq, result)
}

/*
*  ======================== levelOrder 1D Array ========================
*/

func levelOrderIterative(node *TreeNode) []int {
    queue := list.New()
    queue.PushBack(node)
    arr := []int{}
    for ; queue.Len() != 0; {
        front := queue.Front()
        curr := queue.Remove(front).(*TreeNode)
        arr = append(arr, curr.Val)
        if curr.Left != nil {
            queue.PushBack(curr.Left)
        }
        if curr.Right != nil {
            queue.PushBack(curr.Right)
        }
    }
    return arr
}

// Construct Tree from Array
// pass in range from 0 to array length
// in recursive calc the pos of root node
//     check if start < pos, pass in start, pos
//     check if pos+1 < end, pass in pos+1, end
//     constructTreeRecursive(root, nums, 0, len(nums))
func constructTreeRecursive(node *TreeNode, nums []int, start, end int) {
    // calculate the position of the root node
    pos := (end - start)/2
    node.Val = nums[pos]
    if start < pos {
        node.Left = &TreeNode{}
        constructTreeRecursive(node.Left, nums, start, pos)
    }
    if pos+1 < end {
        node.Right = &TreeNode{}
        constructTreeRecursive(node.Right, nums, pos+1, end)
    }
}

func constructTreeFromArray(nums []int) *TreeNode {
    return constructTreeFromArrayRecursive(nil, nums, 0)
}

func constructTreeFromArrayRecursive(node *TreeNode, nums []int, pos int) *TreeNode {
    if pos >= len(nums) {
        return nil
    }
    if nums[pos] == 0 {
        return nil
    }
    if node == nil {
        node = &TreeNode{}
    }
    node.Val = nums[pos]
    node.Left = constructTreeFromArrayRecursive(node.Left, nums, pos*2+1)
    node.Right = constructTreeFromArrayRecursive(node.Right, nums, pos*2+2)
    return node
}
