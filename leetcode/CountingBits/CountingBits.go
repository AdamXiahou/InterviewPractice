package CountingBits

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