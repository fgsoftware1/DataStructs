package main

import (
	"fmt"
	"data"
)

func main() {
	// Stack testing
	stack := data.NewStack()
	
	// Test pushing elements
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	fmt.Println("Stack after pushing:", stack.Items())
	
	// Test popping elements
	val1, _ := stack.Pop()
	fmt.Println("Popped:", val1)
	val2, _ := stack.Pop()
	fmt.Println("Popped:", val2)
	
	// Test isEmpty
	fmt.Println("Is stack empty?", stack.IsEmpty())

	// Test peek
	lastElement, _ := stack.Peek()
	fmt.Println("Last element in stack:", lastElement)
	
	// Pop remaining element
	stack.Pop()
	fmt.Println("Is stack empty after popping all?", stack.IsEmpty())

	// Queue testing
	queue := data.NewQueue()

	// Test enqueueing elements
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println("\nQueue after enqueueing:", queue.Items())

	// Test dequeueing elements
	val1 = queue.Dequeue().(int)
	fmt.Println("Dequeued:", val1)
	val2 = queue.Dequeue().(int)
	fmt.Println("Dequeued:", val2)

	// Test isEmpty
	fmt.Println("Is queue empty?", queue.IsEmpty())

	// Test peek
	firstElement := queue.Peek()
	fmt.Println("First element in queue:", firstElement)

	// Dequeue remaining element
	queue.Dequeue()
	fmt.Println("Is queue empty after dequeuing all?", queue.IsEmpty())

	// Linked List testing
	list := data.NewLinkedList()

	// Test Inserting elements
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)

	// Test traversing the list
	fmt.Println("\nList after inserting:")
	list.Traverse()

	// Test removing elements
	list.Remove(2)
	fmt.Println("List after removing 2:", list.Items())

	// Test contains
	fmt.Println("List contains 2?", list.Contains(2))
	fmt.Println("List contains 1?", list.Contains(1))

	// Test size
	fmt.Println("List size:", list.Size())

	// Test clear
	list.Clear()
	fmt.Println("Is list empty after clearing?", list.Size() == 0)
}