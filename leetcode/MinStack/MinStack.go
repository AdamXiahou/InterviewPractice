package MinStack

import "math"

type MinStack struct {
    stack []int
    minStack []int
}

func Constructor() MinStack {
    return MinStack{
        stack: []int{},
        minStack: []int{math.MaxInt64},   //math.MaxInt64 => Untyped Constant 代表任意精度的值并且不会溢出   
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