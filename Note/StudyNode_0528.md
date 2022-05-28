# 数据库视频进度 
看到P40 instead of 触发器              
# 数据库练习题         
看了五道题（共看了40道）                          
# Leetcode算法
###  反转字符串
题目：           
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。         
不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。         
思路： 双指针  左0 ，右最大 左++ 右--         
func reverseString(s []byte) {         
    for left, right := 0, len(s)-1; left < right; left++ {         
        s[left], s[right] = s[right], s[left]         
        right--         
    }         
}         
         
### 不同的二叉搜索树 II
题目：给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。         
思路：回溯         
Definition for a binary tree node.         
type TreeNode struct {         
      Val int         
      Left *TreeNode         
      Right *TreeNode         
}         
         
func generateTrees(n int) []*TreeNode {         
    if n == 0 {         
        return nil         
    }         
    return helper(1, n)         
}         
         
func helper(start, end int) []*TreeNode {         
    if start > end {         
        return []*TreeNode{nil}         
    }         
    allTrees := []*TreeNode{}         
    // 枚举可行根节点         
    for i := start; i <= end; i++ {         
        // 获得所有可行的左子树集合         
        leftTrees := helper(start, i - 1)         
        // 获得所有可行的右子树集合         
        rightTrees := helper(i + 1, end)         
        // 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上         
        for _, left := range leftTrees {         
            for _, right := range rightTrees {         
                currTree := &TreeNode{i, nil, nil}         
                currTree.Left = left         
                currTree.Right = right         
                allTrees = append(allTrees, currTree)         
            }         
        }         
    }         
    return allTrees         
}         
         
