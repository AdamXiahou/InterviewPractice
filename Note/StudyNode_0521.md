# 数据库视频进度
看到P7 B+树     
# 数据库练习题
看了十道题    
# Leetcode算法

### 二叉树的最大深度
题目： 给定一个二叉树，找出其最大深度   
思路： 找左右子树最大深度，然后加一     
func maxDepth(root *TreeNode) int {     
    if root == nil {    
        return 0    
    }    
    return max(maxDepth(root.Left), maxDepth(root.Right)) + 1    
}    
    
func max(a, b int) int {    
    if a > b {    
        return a    
    }    
    return b    
}    
### 平衡二叉树
题目： 给定一个二叉树，判断它是否是高度平衡的二叉树   
思路： 自上而下的递归判断每个节点是否为平衡二叉树     
func isBalanced(root *TreeNode) bool {    
    if root == nil {    
        return true    
    }    
    return abs(height(root.Left) - height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)    
}    
// _return abs(height(root.Left) - height(root.Right)) <= 1_确保当前根节点是平衡二叉树      
// _isBalanced(root.Left)_ 左递归      
//_isBalanced(root.Right)_  右递归      
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
