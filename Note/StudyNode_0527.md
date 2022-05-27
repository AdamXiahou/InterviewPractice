# 数据库视频进度 
看到P40 instead of 触发器              
# 数据库练习题         
看了五道题（共看了40道）                          
# Leetcode算法
### 三角形最小路径和 **
题目：
给定一个三角形 triangle ，找出自顶向下的最小路径和。             
                          
每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。             
             
示例 1：             
             
输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]             
输出：11             
解释：如下面简图所示：             
   2             
  3 4             
 6 5 7             
4 1 8 3             
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。             
             
思路：
方法一：动态规划  //其中存储路径的辅助数组f[]是三角矩阵，可以压缩存储节省空间复杂度
func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    f := make([][]int, n)
    for i := 0; i < n; i++ {
        f[i] = make([]int, n)       //初始化F存储最小路径和
    }
    f[0][0] = triangle[0][0]
    for i := 1; i < n; i++ {
        f[i][0] = f[i - 1][0] + triangle[i][0]                  //左边直接加下来
        for j := 1; j < i; j++ {
            f[i][j] = min(f[i - 1][j - 1], f[i - 1][j]) + triangle[i][j]       //取相邻最小值分别与对应的triangle[][]相加
        }
        f[i][i] = f[i - 1][i - 1] + triangle[i][i]              //右边直接加下来
    }
    ans := math.MaxInt32   //定义ans
    for i := 0; i < n; i++ {
        ans = min(ans, f[n-1][i])   //取所有路径和最小值
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

###  爬楼梯
题目：             
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。             
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？             
思路：
方法一：动态规划  // 方法数位 1 2 3 5 8 相当于后一个数等于前两个数相加
func climbStairs(n int) int {
    p, q, r := 0, 0, 1
    for i := 1; i <= n; i++ {
        p = q
        q = r
        r = p + q
    }
    return r
}
