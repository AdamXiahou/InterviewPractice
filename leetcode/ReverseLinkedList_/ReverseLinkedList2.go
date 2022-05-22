package ReverseLinkedList_

//Definition for singly-linked list.           
// type ListNode struct {           
// 	Val int           
// 	Next *ListNode      
// }        
func reverseList2(head *ListNode) *ListNode {     
  if head == nil || head.Next == nil {     
	  return head     
  }     
  newHead := reverseList2(head.Next)     
  head.Next.Next = head     
  head.Next = nil     
  return newHead     
}     
 