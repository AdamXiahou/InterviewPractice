package CloneGraph_

type Node1 struct {                           
    Val int                           
    Neighbors []*Node1                           
}                      
func cloneGraph1(node *Node1) *Node1 {           
    if node == nil {           
        return node           
    }           
    visited := map[*Node1]*Node1{}           
           
    // 将题目给定的节点添加到队列           
    queue := []*Node1{node}           
    // 克隆第一个节点并存储到哈希表中           
    visited[node] = &Node1{node.Val, []*Node1{}}           
           
    // 广度优先搜索           
    for len(queue) > 0 {           
        // 取出队列的头节点           
        n := queue[0]           
        // 遍历该节点的邻居           
        queue = queue[1:]           
        for _, neighbor := range n.Neighbors {                      
            if _, ok := visited[neighbor]; !ok {           
                // 如果没有被访问过，就克隆并存储在哈希表中           
                visited[neighbor] = &Node1{neighbor.Val, []*Node1{}}           
                // 将邻居节点加入队列中           
                queue = append(queue, neighbor)           
            }           
            // 更新当前节点的邻居列表           
            visited[n].Neighbors = append(visited[n].Neighbors, visited[neighbor])           
        }           
    }           
    return visited[node]           
}           
           