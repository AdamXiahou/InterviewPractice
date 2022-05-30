# C#视频进度 
https://www.youtube.com/watch?v=GhQdlIFylQ8&list=LL&index=1&t=231s(Working with numbers)                                
# Leetcode算法
### 二叉搜索树中的插入操作
题目：         
给定二叉搜索树（BST）的根节点 root 和要插入树中的值 value ，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。         
         
注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。   
               
思路：//从根节点开始根据二叉树左小右大定理依次往下寻找          
// Definition for a binary tree node.         
type TreeNode struct {                  
    Val int         
    Left *TreeNode         
    Right *TreeNode         
}         
func insertIntoBST(root *TreeNode, val int) *TreeNode {         
    if root == nil {         
        return &TreeNode{Val: val}         
    }         
    p := root         
    for p != nil {         
        if val < p.Val {         
            if p.Left == nil {         
                p.Left = &TreeNode{Val: val}         
                break         
            }         
            p = p.Left         
        } else {         
            if p.Right == nil {         
                p.Right = &TreeNode{Val: val}         
                break         
            }         
            p = p.Right         
        }         
    }         
    return root         
}         
         
### 删除二叉搜索树中的节点
题目：       
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。      
      
一般来说，删除节点可分为两个步骤：      
      
1. 首先找到需要删除的节点；      
2. 如果找到了，删除它。      
      
思路：      
// Definition for a binary tree node.         
type TreeNode struct {                  
    Val int         
    Left *TreeNode         
    Right *TreeNode         
}      
// 递归
func deleteNode(root *TreeNode, key int) *TreeNode {      
    if root==nil{      
        return nil      
    }      
    if key<root.Val{      
        root.Left=deleteNode(root.Left,key)      
        return root      
    }      
    if key>root.Val{      
        root.Right=deleteNode(root.Right,key)      
        return root      
    }      
    if root.Right==nil{      
        return root.Left      
    }      
    if root.Left==nil{      
        return root.Right      
    }      
    minnode:=root.Right      
    for minnode.Left!=nil{      
        minnode=minnode.Left      
    }      
    root.Val=minnode.Val      
    root.Right=deleteNode1(root.Right)      
    return root      
}      
      
func deleteNode1(root *TreeNode)*TreeNode{      
    if root.Left==nil{      
        pRight:=root.Right      
        root.Right=nil      
        return pRight      
    }      
    root.Left=deleteNode1(root.Left)      
    return root      
}            
