# Leetcode算法
### 对角线遍历
题目：                
给你一个大小为 m x n 的矩阵 mat ，请以对角线遍历的顺序，用一个数组返回这个矩阵中的所有元素。               
                
思路：//对角线个数为 m+n-1;区分奇偶，从i=0开始偶数就从下往上x++，y--。奇数就从上往下x--,y++。依次输入数组ans               
func findDiagonalOrder(mat [][]int) []int {               
    m, n := len(mat), len(mat[0])               
    ans := make([]int, 0, m*n)               
    for i := 0; i < m+n-1; i++ {               
        if i%2 == 1 {                           
            x := max(i-n+1, 0)               
            y := min(i, n-1)               
            for x < m && y >= 0 {               
                ans = append(ans, mat[x][y])               
                x++               
                y--               
            }               
        } else {               
            x := min(i, m-1)               
            y := max(i-m+1, 0)               
            for x >= 0 && y < n {               
                ans = append(ans, mat[x][y])               
                x--               
                y++               
            }                 
        }               
    }               
    return ans               
}               
               
func min(a, b int) int {               
    if a > b {               
        return b               
    }               
    return a               
}               
               
func max(a, b int) int {               
    if b > a {               
        return b               
    }               
    return a               
}               
               
### 最长公共前缀
题目：               
编写一个函数来查找字符串数组中的最长公共前缀。               
如果不存在公共前缀，返回空字符串 ""。               
               
思路：               
func longestCommonPrefix(strs []string) string {               
    if len(strs) == 0 {               
        return ""               
    }               
    prefix := strs[0]               
    count := len(strs)                              
    for i := 1; i < count; i++ {               
        prefix = lcp(prefix, strs[i])  //直接比，比完prefix就是新的最长公共前缀，一直遍历到最后一个字符串               
        if len(prefix) == 0 {               
            break               
        }               
    }               
    return prefix               
}               
               
func lcp(str1, str2 string) string {               
    length := min(len(str1), len(str2))               
    index := 0               
    for index < length && str1[index] == str2[index] {               
        index++               
    }               
    return str1[:index]               
}               
               
func min(x, y int) int {               
    if x < y {               
        return x               
    }               
    return y               
}               
               