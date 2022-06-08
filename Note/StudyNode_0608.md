# Leetcode算法
### 斐波那契数
题目：          
斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：          
          
F(0) = 0，F(1) = 1          
F(n) = F(n - 1) + F(n - 2)，其中 n > 1          
给定 n ，请计算 F(n)           
               
思路：//动态规划，小于二时等于他自己，剩下的每个等于前两个之和，和昨天的路径和相似          
func fib(n int) int {          
    if n < 2 {          
        return n          
    }          
    p, q, r := 0, 0, 1          
    for i := 2; i <= n; i++ {          
        p = q          
        q = r          
        r = p + q                    
    }          
    return r          
}          
          
### 找到字符串中所有字母异位词          
题目：        
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。        
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。        
        
思路：        
func findAnagrams(s, p string) (ans []int) {        
    sLen, pLen := len(s), len(p)        
    if sLen < pLen {        
        return        
    }        
        
    var sCount, pCount [26]int        
    for i, ch := range p {        
        sCount[s[i]-'a']++        
        pCount[ch-'a']++        
    }        
    if sCount == pCount {        
        ans = append(ans, 0)        
    }        
        
    for i, ch := range s[:sLen-pLen] { //滑动窗口根据字符串长度依次遍历S的子序列        
        sCount[ch-'a']--        
        sCount[s[i+pLen]-'a']++        
        if sCount == pCount {        
            ans = append(ans, i+1)        
        }        
    }        
    return        
}        
        