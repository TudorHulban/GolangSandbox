package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Data     int
	nextNode *Node
}

type LinkedList struct {
	Head *Node
}

var ErrorNotFound = errors.New("value was not found")

func NewLinkedList(n *Node) *LinkedList {
	return &LinkedList{
		Head: n,
	}
}

func (l *LinkedList) Prepend(n *Node) {
	n.nextNode = l.Head
	l.Head = n
}

func (l *LinkedList) DeleteValue(v int) error {
	if l.Head == nil {
		return errors.New("empty list")
	}

	if l.Head.Data == v {
		if l.Head.nextNode != nil {
			l.Head = l.Head.nextNode
			return nil
		}

		l.Head = nil
		return errors.New("empty list")
	}

	// if previous did not trigger the one value list does not have the value.
	if l.Head.nextNode == nil {
		return ErrorNotFound
	}

	next := l.Head

	for next.nextNode.Data != v {
		if next.nextNode.nextNode == nil {
			return ErrorNotFound
		}

		if next.nextNode == nil {
			return errors.New("stopping , only head remains")
		}
		next = next.nextNode
	}

	next.nextNode = next.nextNode.nextNode
	return nil
}

func (l LinkedList) PrintData() {
	next := l.Head

	for next != nil {
		fmt.Printf("%d ", next.Data)
		next = next.nextNode
	}

	fmt.Print("\n")
}

func main() {
	n1 := Node{
		Data: 1,
	}

	l := NewLinkedList(&n1)

	n2 := Node{
		Data: 2,
	}

	n3 := Node{
		Data: 3,
	}

	l.Prepend(&n2)
	l.Prepend(&n3)

	fmt.Println(*l.Head)
	l.PrintData()

	l.DeleteValue(3)
	l.PrintData()

	l.DeleteValue(2)
	l.PrintData()

	fmt.Println(l.DeleteValue(4))

	l.DeleteValue(1)
	l.PrintData()

	fmt.Println(l.DeleteValue(4))
}
