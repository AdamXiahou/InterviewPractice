# C#视频进度 
https://www.youtube.com/watch?v=GhQdlIFylQ8&list=LL&index=1&t=231s(Working with numbers)                                
# Leetcode算法
### 二叉树的最近公共祖先
题目：            
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。            
最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”            
思路：            //明确何为最近公共祖先，再来依次递归二叉树         
//Definition for a binary tree node.            
type TreeNode struct {            
    Val int            
    Left *TreeNode            
    Right *TreeNode            
}            
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {            
    if root == nil {            
        return nil            
    }            
    if root.Val == p.Val || root.Val == q.Val {            
        return root            
    }            
    left := lowestCommonAncestor(root.Left, p, q)            
    right := lowestCommonAncestor(root.Right, p, q)            
    if left != nil && right != nil {            
        return root            
    }            
    if left == nil {            
        return right            
    }            
    return left            
}            
            
