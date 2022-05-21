package AddStrings

import "strconv"

func addStrings(num1 string, num2 string) string {       
    add := 0    //进位        
    ans := ""    //新字符串      
    for i, j := len(num1) - 1, len(num2) - 1; i >= 0 || j >= 0 || add!= 0; i, j = i - 1, j - 1 {       
        var x, y int      
        if i >= 0 {      
            x = int(num1[i] - '0')     
        }      
        if j >= 0 {      
            y = int(num2[j] - '0')      
        }      
        result := x + y + add       
        ans = strconv.Itoa(result%10) + ans    //strconv.Itoa 把整型转成字符串并确保是如果有进位 新字符串各位数            
        add = result / 10     
    }       
    return ans       
}        
