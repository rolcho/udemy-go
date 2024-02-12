package main

import (
	"fmt"

	"example.com/arrays/stack"
)

func testStack() {
	var myStack stack.Stack

	fmt.Println(myStack.IsEmpty())

	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Push(4)

	for !myStack.IsEmpty() {

		popped, popError := myStack.Pop()
		if popError != nil {
			fmt.Println("ERROR:", popError)
		} else {
			fmt.Println("Popped:", popped)
		}

		lastItem, lastError := myStack.Peek()
		if lastError != nil {
			fmt.Println("ERROR:", lastError)
		} else {
			fmt.Println("Last item:", lastItem)
		}

		fmt.Println("---------------------")
	}
}
