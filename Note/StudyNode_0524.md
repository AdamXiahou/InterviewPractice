# 数据库视频进度
看到P21 幻读           
# 数据库练习题
看了十道题（共看了25道）        
# Leetcode算法
### 最小栈
题目：      
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。              
实现 MinStack 类:      
MinStack() 初始化堆栈对象。      
void push(int val) 将元素val推入堆栈。      
void pop() 删除堆栈顶部的元素。      
int top() 获取堆栈顶部的元素。      
int getMin() 获取堆栈中的最小元素。      
思路：       

type MinStack struct {      
    stack []int      
    minStack []int      
}      
      
func Constructor() MinStack {      
    return MinStack{      
        stack: []int{},      
        minStack: []int{math.MaxInt64}, //math.MaxInt64 => Untyped Constant 代表任意精度的值并且不会溢出         
    }      
}      
      
func (this *MinStack) Push(x int)  {      
    this.stack = append(this.stack, x)          
    top := this.minStack[len(this.minStack)-1]      
    this.minStack = append(this.minStack, min(x, top))     //把每次push对应的min值存入minStack中      
}      
      
func (this *MinStack) Pop()  {      
    this.stack = this.stack[:len(this.stack)-1]      
    this.minStack = this.minStack[:len(this.minStack)-1]    //在切片中表示 :len(this.minStack)-1 减去最大长度减一的那个切片比如原来Stack 的len是10 现在的len则为9      
}      
      
func (this *MinStack) Top() int {      
    return this.stack[len(this.stack)-1]      
}      
      
func (this *MinStack) GetMin() int {      
    return this.minStack[len(this.minStack)-1]        
}      
      
func min(x, y int) int {      
    if x < y {      
        return x      
    }      
    return y      
}      
      
### 字符串解码 *      
题目：         
给定一个经过编码的字符串，返回它解码后的字符串。           
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。       
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。      
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。      
思路：       
方法一：栈操作      //利用栈来辅助完成字符串的解码       
func decodeString(s string) string {      
    stk := []string{}      
    ptr := 0      
    for ptr < len(s) {      
        cur := s[ptr]      
        if cur >= '0' && cur <= '9' {      
            digits := getDigits(s, &ptr)      
            stk = append(stk, digits)      
        } else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {      
            stk = append(stk, string(cur))      
            ptr++      
        } else {      
            ptr++      
            sub := []string{}      
            for stk[len(stk)-1] != "[" {      
                sub = append(sub, stk[len(stk)-1])      
                stk = stk[:len(stk)-1]      
            }      
            for i := 0; i < len(sub)/2; i++ {      
                sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]      
            }      
            stk = stk[:len(stk)-1]      
            repTime, _ := strconv.Atoi(stk[len(stk)-1])      //strconv.Atoi (str string) (int, error)将字符串类型转换为int类型       
            stk = stk[:len(stk)-1]      
            t := strings.Repeat(getString(sub), repTime)      //解码为重复字符串       
            stk = append(stk, t)      
        }      
    }      
    return getString(stk)      
}      
      
func getDigits(s string, ptr *int) string {      
    ret := ""      
    for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {      
        ret += string(s[*ptr])      
    }      
    return ret      
}      
      
func getString(v []string) string {      
    ret := ""      
    for _, s := range v {      
        ret += s      
    }      
    return ret      
}      
      
方法二：递归            //从左往右依次遍历如果有数字就跳过[进入递归       
var (      
    src string      
    ptr int      
)      
      
func decodeString(s string) string {      
    src = s      
    ptr = 0      
    return getString()      
}      
      
func getString() string {      
    if ptr == len(src) || src[ptr] == ']' {      
        return ""      
    }      
    cur := src[ptr]      
    repTime := 1      
    ret := ""      
    if cur >= '0' && cur <= '9' {      //判断数字       
        repTime = getDigits()      
        ptr++      
        str := getString()      
        ptr++      
        ret = strings.Repeat(str, repTime)      
    } else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' {      
        ret = string(cur)      
        ptr++      
    }      
    return ret + getString()      
}      
      
func getDigits() int {      
    ret := 0      
    for ; src[ptr] >= '0' && src[ptr] <= '9'; ptr++ {      
        ret = ret * 10 + int(src[ptr] - '0')      
    }      
    return ret      
}      