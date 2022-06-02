# C#视频进度 
https://www.youtube.com/watch?v=GhQdlIFylQ8&list=LL&index=1&t=231s (If Statements)                                
# Leetcode算法
###  逆波兰表达式求值
题目：                                
根据 逆波兰表示法，求表达式的值。                                
有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。                                
注意 两个整数之间的除法只保留整数部分。                                
可以保证给定的逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况                                
                                
思路：//创建一个栈存储表达式,如果是数值就存储，如果是表达式就把最近两个数值提出与表达式计算后再存储                                
func evalRPN(tokens []string) int {                                
    stack := []int{}                                
    for _, token := range tokens {                                
        val, err := strconv.Atoi(token)                                
        if err == nil {                                
            stack = append(stack, val)                                
        } else {                                
            num1, num2 := stack[len(stack)-2], stack[len(stack)-1]                                
            stack = stack[:len(stack)-2]                                
            switch token {                                
            case "+":                                
                stack = append(stack, num1+num2)                                
            case "-":                                
                stack = append(stack, num1-num2)                                
            case "*":                
                stack = append(stack, num1*num2)                
            default:                
                stack = append(stack, num1/num2)                
            }                
        }                
    }                
    return stack[0]                
}                
###  克隆图
题目：                                
给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。                
图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。                
class Node {                
    public int val;                
    public List<Node> neighbors;                
}                
思路：  //深度优先和广度优先遍历   
一： 深度优先        
//Definition for a Node.                
type Node struct {                
    Val int                
    Neighbors []*Node                
}                
func cloneGraph(node *Node) *Node {           
    visited := map[*Node]*Node{}                      
    var cg func(node *Node) *Node           
    cg = func(node *Node) *Node {           
        if node == nil {           
            return node           
        }           
           
        // 如果该节点已经被访问过了，则直接从哈希表中取出对应的克隆节点返回           
        if _, ok := visited[node]; ok {           
            return visited[node]           
        }           
           
        // 克隆节点，注意到为了深拷贝我们不会克隆它的邻居的列表           
        cloneNode := &Node{node.Val, []*Node{}}           
        // 哈希表存储           
        visited[node] = cloneNode           
           
        // 遍历该节点的邻居并更新克隆节点的邻居列表           
        for _, n := range node.Neighbors {           
            cloneNode.Neighbors = append(cloneNode.Neighbors, cg(n))           
        }           
        return cloneNode           
    }           
    return cg(node)           
}           
二： 广度优先           
type Node struct {                           
    Val int                           
    Neighbors []*Node                           
}                      
func cloneGraph(node *Node) *Node {           
    if node == nil {           
        return node           
    }           
    visited := map[*Node]*Node{}           
           
    // 将题目给定的节点添加到队列           
    queue := []*Node{node}           
    // 克隆第一个节点并存储到哈希表中           
    visited[node] = &Node{node.Val, []*Node{}}           
           
    // 广度优先搜索           
    for len(queue) > 0 {           
        // 取出队列的头节点           
        n := queue[0]           
        // 遍历该节点的邻居           
        queue = queue[1:]           
        for _, neighbor := range n.Neighbors {                      
            if _, ok := visited[neighbor]; !ok {           
                // 如果没有被访问过，就克隆并存储在哈希表中           
                visited[neighbor] = &Node{neighbor.Val, []*Node{}}           
                // 将邻居节点加入队列中           
                queue = append(queue, neighbor)           
            }           
            // 更新当前节点的邻居列表           
            visited[n].Neighbors = append(visited[n].Neighbors, visited[neighbor])           
        }           
    }           
    return visited[node]           
}           
           