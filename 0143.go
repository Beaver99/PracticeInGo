package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	// 找到链表的右半部分，slow是右子链表的头
	slow, fast := head, head
	var tmp *ListNode
	//for fast != nil&&fast.Next.Next != nil {
	//	slow = slow.Next
	//	fast = fast.Next.Next
	//}
	//if fast != nil { // 奇数节点要再后移一个
	//	slow = slow.Next
	//}

	for fast != nil{
		if fast.Next == nil || fast.Next.Next == nil{// 奇数/偶数
			tmp = slow.Next
			slow.Next = nil//断开右子链
			slow = tmp
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 倒置右链表
	rHead := slow
	nxt := slow.Next
	for nxt != nil {
		tmp = nxt.Next
		nxt.Next = rHead

		slow.Next = tmp

		rHead = nxt

		nxt = tmp
	}

	//fmt.Println("hhh")
	// 合并链表
	rCurr, curr := rHead, head
	var rTmp *ListNode
	for rCurr != nil {
		rTmp = rCurr.Next
		tmp = curr.Next

		curr.Next = rCurr

		rCurr.Next = tmp

		rCurr = rTmp
		curr = tmp
	}
}

func main() {
	var n1, n2, n3, n4, n5,n6,n7 ListNode
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5
	n5.Next = &n6
	n6.Next = &n7

	//reorderList(&n1)

	for tmp, i := &n1, 1; tmp != nil; tmp, i = tmp.Next, i+1 {
		tmp.Val = i
	}
	reorderList(&n1)

	for tmp:= &n1; tmp != nil; tmp= tmp.Next {
		fmt.Println(*tmp)
	}

	//fmt.Println("hahaha")
}
