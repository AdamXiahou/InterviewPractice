# 数据库视频进度 
看到P28 T-SQL编程命名规则              
# 数据库练习题         
看了五道题（共看了30道）                 
# Leetcode算法
### 只出现一次的数字
题目： 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。         
思路：可以用集合或者哈希表方式做辅助，重新存储数字对比，或者用位运算，利用异或的交换结合律         
func singleNumber(nums []int) int {         
    single := 0         
    for _, num := range nums { //range 关键字用于 迭代数组(array)、切片(slice)、链表(channel)或集合(map)的元素         
        single ^= num // **^运算 single = single^num  二元运算符就是异或(1^1 =0, 0^0=0,1^0=1,0^1=1)   ^num 一元运算符表示是按位取反**         
    }         
    return single         
}         

###  位1的个数 *
题目： 编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。         
思路：         
方法一：循环检查二进制位         
func hammingWeight(num uint32) (ones int) {         
    for i := 0; i < 32; i++ {         
        if 1<<i&num > 0 {           
            ones++         
        }         
    }         
    return         
}         
方法二：位运算优化
func hammingWeight(num uint32) (ones int) {         
    for ; num > 0; num &= num - 1 {         
        ones++         
    }         
    return         
}         