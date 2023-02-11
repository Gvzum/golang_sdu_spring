package stack

import (
	"errors"
	"fmt"
)

type Node struct {
	Value   int
	pointer *Node
}

type Stack struct {
	last_node *Node
}

func (s *Stack) Pop() (int, error) {
	fmt.Println("Pop element from")

	if s.last_node == nil {
		return 0, errors.New("Doesn't have element in stack error.")
	}

	next := s.last_node.pointer
	return_value := s.last_node.Value
	s.last_node = nil
	s.last_node = next

	return return_value, nil
}

func (s *Stack) Peek() (int, error) {
	fmt.Println("Peek method.")

	if s.last_node == nil {
		return 0, errors.New("Doesn't have element in stack error.")
	}
	return s.last_node.Value, nil
}

func (s *Stack) Push(node *Node) {
	fmt.Println("Push method.")

	if s.last_node == nil {
		s.last_node = node
	} else {
		tmp := s.last_node
		s.last_node = node
		s.last_node.pointer = tmp
	}
}

func (s *Stack) Clear() {
	fmt.Println("Clear method.")

	for s.last_node != nil {
		tmp := s.last_node
		s.last_node = nil
		s.last_node = tmp.pointer
	}
}

func (s *Stack) Contains(target int) bool {
	fmt.Println("Contains method.")

	node := s.last_node
	for node != nil {
		if node.Value == target {
			return true
		}
	}
	return false
}

func (s *Stack) Increment() {
	fmt.Println("Increment method.")

	node := s.last_node
	for node != nil {
		node.Value++
		node = node.pointer
	}
}

func (s *Stack) Print() {
	fmt.Println("Print method.")

	node := s.last_node
	for node != nil {
		fmt.Println(node.Value)
		node = node.pointer
	}
}

func (s *Stack) PrintReverse() {
	fmt.Println("PrintReverse method.")

	list_of_node := []int{}
	node := s.last_node
	for node != nil {
		list_of_node = append(list_of_node, node.Value)
		node = node.pointer
	}

	for i := len(list_of_node) - 1; i >= 0; i-- {
		fmt.Println(list_of_node[i])
	}
}
