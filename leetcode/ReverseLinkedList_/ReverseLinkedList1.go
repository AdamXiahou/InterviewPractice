package ReverseLinkedList_

//Definition for singly-linked list.           
type ListNode struct {      
	Val int      
	Next *ListNode      
}      
func reverseList1(head *ListNode) *ListNode {     
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
