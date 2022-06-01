# C#视频进度 
https://www.youtube.com/watch?v=GhQdlIFylQ8&list=LL&index=1&t=231s (If Statements)                                
# Leetcode算法
### 分隔链表
题目：                 
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。                 
你应当 保留 两个分区中每个节点的初始相对位置。                 
                 
思路： 新建两个辅助链表分别存储大于x和小于x的值，然后合并                 
//Definition for singly-linked list.                 
type ListNode struct {                 
    Val int                 
    Next *ListNode                 
}                 
func partition(head *ListNode, x int) *ListNode {                 
    small := &ListNode{}                 
    smallHead := small                 
    large := &ListNode{}                 
    largeHead := large                 
    for head != nil {                 
        if head.Val < x {                 
            small.Next = head                 
            small = small.Next                 
        } else {                                  
            large.Next = head                 
            large = large.Next                 
        }                 
        head = head.Next                 
    }                 
    large.Next = nil                 
    small.Next = largeHead.Next                 
    return smallHead.Next                 
}                 
                 
### 排序链表
题目：给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。                 
                 
思路：归并排序                 
//Definition for singly-linked list.                 
type ListNode struct {                 
    Val int                 
    Next *ListNode                 
}                 
func merge(head1, head2 *ListNode) *ListNode {             //把给定节点按升序排序          
    dummyHead := &ListNode{}                 
    temp, temp1, temp2 := dummyHead, head1, head2                 
    for temp1 != nil && temp2 != nil {                 
        if temp1.Val <= temp2.Val {                 
            temp.Next = temp1                 
            temp1 = temp1.Next                 
        } else {                 
            temp.Next = temp2                 
            temp2 = temp2.Next                 
        }                 
        temp = temp.Next                 
    }                 
    if temp1 != nil {                 
        temp.Next = temp1                 
    } else if temp2 != nil {                 
        temp.Next = temp2                 
    }                 
    return dummyHead.Next                 
}                 
                 
func sort(head, tail *ListNode) *ListNode {            //拆分链表并归             
    if head == nil {                 
        return head                 
    }                 
                 
    if head.Next == tail {                 
        head.Next = nil                 
        return head                 
    }                 
                 
    slow, fast := head, head                 
    for fast != tail {                 
        slow = slow.Next                 
        fast = fast.Next                 
        if fast != tail {                 
            fast = fast.Next                 
        }                 
    }                 
                 
    mid := slow                 
    return merge(sort(head, mid), sort(mid, tail))                 
}                 
                 
func sortList(head *ListNode) *ListNode {                 
    return sort(head, nil)                 
}                 
                 