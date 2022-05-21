package BalancedBinaryTree

//Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {    
    if root == nil {    
        return true    
    }    
    return abs(height(root.Left) - height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)    
}    
// _return abs(height(root.Left) - height(root.Right)) <= 1_确保当前根节点是平衡二叉树      
// _isBalanced(root.Left)_ 左递归      
// _isBalanced(root.Right)_  右递归      
//比较二叉树最大深度算法    
func height(root *TreeNode) int {    
    if root == nil {    
        return 0    
    }    
    return max(height(root.Left), height(root.Right)) + 1    
}    
   
func max(x, y int) int {    
    if x > y {    
        return x    
    }    
    return y    
}    
//防止深度差为负数     
func abs(x int) int {      
    if x < 0 {    
        return -1 * x     
    }      
    return x     
}      
