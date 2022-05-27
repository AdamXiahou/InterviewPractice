package ClimbingStairs

// 方法数位 1 2 3 5 8 相当于后一个数等于前两个数相加
func climbStairs(n int) int {
    p, q, r := 0, 0, 1
    for i := 1; i <= n; i++ {
        p = q
        q = r
        r = p + q
    }
    return r
}
