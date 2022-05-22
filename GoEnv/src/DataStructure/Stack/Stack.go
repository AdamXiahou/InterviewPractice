package main
//用数组实现栈
import "fmt"
//定义一个新的类型来表达栈，需要一个数组和一个指向最后一个元素的索引
type Stack struct {
	i int
	data [10]int
}
//注意，Go的数据传递中，是值传递,但是这样不能满足我们的预期。因为一个副本被创建并传递给函数，push和pop处理的都是副本。所以提供指针
func (s *Stack) push(k int) {
	if s.i +1 > 10 {
		fmt.Println("I'm full")
		return
	}
	s.data[s.i] = k
	s.i++
}
func (s *Stack) pop() int{
	if s.i -1 < 0 {
		fmt.Println("I'm hungry")
		return 0
	}
 
	s.i--
	result := s.data[s.i]
	s.data[s.i] = 0
	return result
}

func main() {
	var s Stack
	s.push(1)
	s.push(11)
	s.push(21)
	s.push(31)
	s.push(41)
	s.push(51)
	s.push(61)
	s.push(71)
	s.push(81)
	s.push(91)
	s.push(101)
	s.push(1011)
	fmt.Printf("Stack %v \n", s)
	x := s.pop()
	fmt.Println("Pop x ", x)
	y := s.pop()
	fmt.Println("Pop y ", y)
	z := s.pop()
	fmt.Println("Pop z ", z)
	fmt.Printf("After pop the Stack is %v \n", s)
	s.push(161)
	fmt.Printf("Stack %v \n", s)
}
