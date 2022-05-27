package Triangle

import "math"

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
