package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type Stack struct {
	Head *Node
	Tail *Node
}

func (s *Stack) Push(Val int) {
	newNode := new(Node)
	newNode.Val = Val
	if s.Head == nil {
		s.Head = newNode
		s.Tail = newNode
	} else {
		newNode.Prev = s.Tail
		s.Tail.Next = newNode
		s.Tail = newNode
	}
}

func (s *Stack) Peek() int {
	return s.Tail.Val
}

func (s *Stack) Clear() {
	s.Head = nil
	s.Tail = nil
}

func (s *Stack) Print() {
	cur := s.Head
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
}

func (s *Stack) Contains(Val int) bool {
	cur := s.Head
	for cur != nil {
		if cur.Val == Val {
			return true
		}
		cur = cur.Next
	}
	return false
}

func (s *Stack) Increment() {
	cur := s.Head
	for cur != nil {
		cur.Val += 1
		cur = cur.Next
	}
}

func (s *Stack) PrintReverse() {
	cur := s.Head
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
}

func (s *Stack) Pop() {
	s.Tail = s.Tail.Prev
	s.Tail.Next = nil
}
