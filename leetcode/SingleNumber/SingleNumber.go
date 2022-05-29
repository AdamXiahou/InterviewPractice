package SingleNumber

func singleNumber(nums []int) int {         
    single := 0         
    for _, num := range nums { //range 关键字用于 迭代数组(array)、切片(slice)、链表(channel)或集合(map)的元素         
        single ^= num // **^运算 single = single^num  二元运算符就是异或(1^1 =0, 0^0=0,1^0=1,0^1=1)   ^num 一元运算符表示是按位取反**         
    }         
    return single         
}  