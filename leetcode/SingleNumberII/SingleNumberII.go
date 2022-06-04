package SingleNumberII

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
