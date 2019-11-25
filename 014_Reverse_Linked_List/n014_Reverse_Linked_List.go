package main

import (
	"Solution_For_Jian_Zhi_Offer/util"
)

func reverseList(head *util.ListNode) *util.ListNode {
	var pre *util.ListNode
	pre, cur := nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
