package DecodeString

import "strings"

var (      
    src string      
    ptr int      
)      
      
func decodeString2(s string) string {      
    src = s      
    ptr = 0      
    return getString2()      
}      
      
func getString2() string {      
    if ptr == len(src) || src[ptr] == ']' {      
        return ""      
    }      
    cur := src[ptr]      
    repTime := 1      
    ret := ""      
    if cur >= '0' && cur <= '9' {      //判断数字
        repTime = getDigits2()      
        ptr++      
        str := getString2()      
        ptr++      
        ret = strings.Repeat(str, repTime)      
    } else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' {      
        ret = string(cur)      
        ptr++      
    }      
    return ret + getString2()      
}      
      
func getDigits2() int {      
    ret := 0      
    for ; src[ptr] >= '0' && src[ptr] <= '9'; ptr++ {      
        ret = ret * 10 + int(src[ptr] - '0')      
    }      
    return ret      
}      