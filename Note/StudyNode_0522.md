# 数据库视频进度
看到P14 数据结构的补充说明      
更新SQL.md中     
# 数据库练习题
看了五道题          
# Leetcode算法

### 删除排序链表中的重复元素
题目：给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。      
思路：链表是已排序的链表，如果有重复的就把重复链表的next指针直接指向下下个元素，丢弃中间重复的元素      
//Definition for singly-linked list.      
type ListNode struct {      
      Val int      
      Next *ListNode      
}      
      
func deleteDuplicates(head *ListNode) *ListNode {      
    if head == nil {      
        return nil      
    }      
      
    cur := head      
    for cur.Next != nil {      
        if cur.Val == cur.Next.Val {      
            cur.Next = cur.Next.Next      
        } else {      
            cur = cur.Next      
        }      
    }      
      
    return head      
}      
### 反转链表*
有点不太好理解      
https://leetcode.cn/problems/reverse-linked-list/solution/dong-hua-yan-shi-206-fan-zhuan-lian-biao-by-user74/     
     
题目：给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。     
思路一：迭代 单链表不为零的情况下，从前到后依次反转.next指针     
//Definition for singly-linked list.           
type ListNode struct {      
      Val int      
      Next *ListNode      
}      
func reverseList(head *ListNode) *ListNode {     
    var prev *ListNode     
    curr := head     
    for curr != nil {     
        next := curr.Next     
        curr.Next = prev     
        prev = curr     
        curr = next     
    }     
    return prev     
}     
     
思路二：递归    单链表不为零的情况下，先用newHead := reverseList(head.Next)循环到链表最后，再依次反转.next指针     
//Definition for singly-linked list.           
type ListNode struct {           
      Val int           
      Next *ListNode      
}        
func reverseList(head *ListNode) *ListNode {     
    if head == nil || head.Next == nil {     
        return head     
    }     
    newHead := reverseList(head.Next)     
    head.Next.Next = head     
    head.Next = nil     
    return newHead     
}     
     