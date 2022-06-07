# Leetcode算法
### 不同路径
题目：         
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？

思路：//动态规划，每一格的路径数为上两格路径数之和，依次相加得到最终结果
func uniquePaths(m, n int) int {
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
        dp[i][0] = 1
    }
    for j := 0; j < n; j++ {
        dp[0][j] = 1
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[m-1][n-1]
}

### 最长递增子序列
题目：
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

思路： //根据nums长度，依次增加长度检验递增序列长度，有则加到dp中存储，返回其最大值maxL
func lengthOfLIS(nums []int) int {
    maxL := 0
    dp := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        dp[i] = 1
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] && dp[j] + 1 > dp[i] {
                dp[i] = dp[j] + 1
            }
        }
        if dp[i] > maxL {
            maxL = dp[i]
        }
    }
    return maxL
}
