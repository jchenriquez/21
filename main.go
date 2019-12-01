package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) (res *ListNode) {
	if l1 == nil && l2 == nil {
		return nil
	}
	var curr *ListNode
	for l1 != nil && l2 != nil {
		switch {
		case l1 == nil:
			if curr == nil {
				res = l2
			} else {
				curr.Next = l2
			}
			break
		case l2 == nil:
			if curr == nil {
				res = l1
			} else {
				curr.Next = l1
			}
			break
		default:
			var minNode **ListNode
			if l1.Val < l2.Val {
				minNode = &l1
			} else {
				minNode = &l2
			}

			if curr == nil {
				curr = *minNode
				res = curr
			} else {
				curr.Next = *minNode
				curr = curr.Next
			}
			*minNode = (*minNode).Next
		}

	}

	return
}

func main() {

}
