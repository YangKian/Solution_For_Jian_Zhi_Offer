package main

import (
	"Solution_For_Jian_Zhi_Offer/util"
	"container/list"
	"fmt"
)

//方法一：双向链表模拟栈
func ReversePrint(lists *util.ListNode) {
	collection := list.New()
	for lists != nil {
		collection.PushBack(lists.Val)
		lists = lists.Next
	}

	size := collection.Len() - 1
	for ; size >= 0; size-- {
		num := collection.Back()
		fmt.Println(num.Value) //注意，collection.Back()返回的是个Element对象，要调用Value方法取出值
		collection.Remove(num)
	}
}

//方法二：递归
func RecursivePrint(lists *util.ListNode) {
	if lists != nil {
		if lists.Next != nil {
			RecursivePrint(lists.Next)
		}
		fmt.Println(lists.Val)
	}
}

func main() {
	a := util.ListNode{Val: 1}
	b := util.ListNode{Val: 2}
	c := util.ListNode{Val: 3}
	d := util.ListNode{Val: 4}
	e := util.ListNode{Val: 5}
	f := util.ListNode{Val: 6}
	g := util.ListNode{Val: 7}
	h := util.ListNode{Val: 8}
	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &e
	e.Next = &f
	f.Next = &g
	g.Next = &h

	ReversePrint(&a)
	fmt.Println("\n---------------------------------------\n")
	RecursivePrint(&a)
}
