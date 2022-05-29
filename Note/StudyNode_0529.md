# 数据库视频进度 
看到P43 触发器的优缺点（教程最后一节）              
# 数据库练习题         
看了十道题（共看了50道）                          
# Leetcode算法
###  字符串的排列 **
题目：                    
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。                    
换句话说，s1 的排列之一是 s2 的 子串 。                    
思路：  使用数组统计s1中各个字符的个数，双指针统计当前遍历的子串中各个字符的个数
func checkInclusion(s1, s2 string) bool {
    n, m := len(s1), len(s2)
    if n > m {
        return false
    }
    cnt := [26]int{}
    for _, ch := range s1 {
        cnt[ch-'a']--
    }
    left := 0
    for right, ch := range s2 {
        x := ch - 'a'
        cnt[x]++
        for cnt[x] > 0 {
            cnt[s2[left]-'a']--
            left++
        }
        if right-left+1 == n {
            return true
        }
    }
    return false
}
