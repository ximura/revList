package list

import (
	"fmt"
	"strings"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func (l *LinkedList[T]) Push(v T) {
	if l.head == nil {
		l.head = &Node[T]{
			Value: v,
		}

		return
	}

	node := l.head
	for node.Next != nil {
		node = node.Next
	}

	node.Next = &Node[T]{
		Value: v,
	}
}

func (l *LinkedList[T]) Reverse() {
	if l.head == nil {
		return
	}

	var prevNode *Node[T]
	node := l.head
	for {
		nextNode := node.Next
		node.Next = prevNode

		prevNode = node

		if nextNode == nil {
			break
		}

		node = nextNode
	}

	l.head = node
}

func (l LinkedList[T]) String() string {
	var builder strings.Builder
	if l.head == nil {
		return ""
	}

	node := l.head
	for node != nil {
		builder.WriteString(fmt.Sprintf("%v", node.Value))
		node = node.Next
		if node != nil {
			builder.WriteString(" -> ")
		}
	}

	return builder.String()
}
