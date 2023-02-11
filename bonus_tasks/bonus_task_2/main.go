package main

import (
	"bonus_task_2/stack"
)

func main() {
	custom_stack := stack.Stack{}
	custom_stack.Push(&stack.Node{Value: 11})
	custom_stack.Push(&stack.Node{Value: 22})
	custom_stack.Push(&stack.Node{Value: 33})
	custom_stack.Push(&stack.Node{Value: 44})
	custom_stack.Push(&stack.Node{Value: 55})
	custom_stack.Push(&stack.Node{Value: 66})
	custom_stack.Print()
	custom_stack.Pop()
	custom_stack.Increment()
	custom_stack.Print()
	custom_stack.Peek()
	custom_stack.PrintReverse()
}
