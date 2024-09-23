package main

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
)

//  2. Add Two Numbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	res := addTwoNumbers(l1, l2)
	fmt.Printf("RES: %v\n", listToSlice(res))
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 342 + 465 = 807
	sl1 := listToSlice(l1)
	sl2 := listToSlice(l2)

	slices.Reverse(sl1)
	slices.Reverse(sl2)

	i1 := sliceToNumber(sl1)
	i2 := sliceToNumber(sl2)

	sl := numberToSlice(i1 + i2)
	slices.Reverse(sl)

	return sliceToList(sl)
}

func listToSlice(l *ListNode) []int {
	var sl []int

	next := l.Next
	sl = append(sl, l.Val)
	for next != nil {
		sl = append(sl, next.Val)
		next = next.Next
	}

	return sl
}

func sliceToNumber(sl1 []int) int {
	var buffer bytes.Buffer
	for _, v := range sl1 {
		buffer.WriteString(strconv.Itoa(v))
	}

	n, _ := strconv.Atoi(buffer.String())

	return n
}

func numberToSlice(i int) []int {
	var sl []int

	for i > 0 {
		sl = append(sl, i%10)
		i = i / 10
	}

	return sl

}

func sliceToList(sl []int) *ListNode {
	if len(sl) == 0 {
		return &ListNode{}
	}

	l := &ListNode{
		Val: sl[0],
	}

	for i := 1; i < len(sl); i++ {
		l = &ListNode{
			Val:  sl[i],
			Next: l,
		}
	}

	return l
}
