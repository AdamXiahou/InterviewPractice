# C#视频进度 
https://www.youtube.com/watch?v=GhQdlIFylQ8&list=LL&index=1&t=231s (Comments)                                
# Leetcode算法
### 比特位计数
题目：
给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。
思路：将x的二进制表示的数与比他小1的数按位与并重新赋值。对x重复该操作，直到x变成0。记录次数
func onesCount(x int) (ones int) {
    for ; x > 0; x &= x - 1 {  //&按位与   1与1则1，其他都为0
        ones++
    }
    return
}

func countBits(n int) []int {
    bits := make([]int, n+1)
    for i := range bits {
        bits[i] = onesCount(i)
    }
    return bits
}
### 只出现一次的数字 II
题目：给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
思路： 
func singleNumber(nums []int) int {
    freq := map[int]int{}   //利用go中自带的Map类型存储出现次数
    for _, v := range nums {
        freq[v]++
    }
    for num, occ := range freq {
        if occ == 1 {
            return num
        }
    }
    return 0 // 不会发生，数据保证有一个元素仅出现一次
}
