package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/add-two-numbers/

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func main() {
	n1 := Node{
		Value: 1,
	}

	l1 := NewLinkedList(&n1)

	n2 := Node{
		Value: 2,
	}

	n3 := Node{
		Value: 1,
	}

	n4 := Node{
		Value: 2,
	}

	l1.Prepend(&n2)
	l1.Prepend(&n3)
	l1.Prepend(&n4)

	number := 2121
	l2 := NewLinkedListFromNumber(number)

	no1, _ := l1.GetNumber()
	no2, _ := l2.GetNumber()

	no := no1 + no2
	fmt.Println(no)
}

func NewLinkedList(n *Node) *LinkedList {
	return &LinkedList{
		Head: n,
	}
}

// NewLinkedListFromNumber Constructor creates linked list from number representation.
// TODO: Add error handling.
func NewLinkedListFromNumber(n int) *LinkedList {
	number := strconv.Itoa(n)

	val, _ := strconv.Atoi(number[len(number)-1:])
	headNode := Node{
		Value: val,
	}

	l := NewLinkedList(&headNode)

	for i := len(number) - 2; i >= 0; i-- {
		v, _ := strconv.Atoi(number[i : i+1])

		l.Prepend(&Node{
			Value: v,
		})
	}

	return l
}

func (l *LinkedList) Prepend(n *Node) {
	n.Next = l.Head
	l.Head = n
}

func (l *LinkedList) GetNumber() (int, error) {
	ints := []int{l.Head.Value}
	next := l.Head

	for next.Next != nil {
		ints = append(ints, next.Next.Value)

		next = next.Next
	}

	var result string

	for j := len(ints) - 1; j >= 0; j-- {
		result = strconv.Itoa(ints[j]) + result
	}

	return strconv.Atoi(result)
}
