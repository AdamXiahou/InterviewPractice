package NumberOneBits_

func hammingWeight2(num uint32) (ones int) {         
    for ; num > 0; num &= num - 1 {         
        ones++         
    }         
    return         
}      