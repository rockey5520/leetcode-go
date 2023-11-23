package main

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	a := addTwoNumbers(
		&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{}}}},
		&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{}}}},
	)
	log.Println("a", a)
	for a != nil {
		log.Println(a.Val)
		a = a.Next
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	curr := res

}
