package data

import "fmt"

type LinkedList struct {
	head *Node
}

type Node struct {
	data interface{}
	next *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil}
}

func (ll *LinkedList) Traverse() {
	current := ll.head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (ll *LinkedList) Insert(data interface{}) {
	newNode := &Node{data: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) Remove(data interface{}) bool {
	if ll.head == nil {
		return false
	}

	if ll.head.data == data {
		ll.head = ll.head.next
		return true
	}

	current := ll.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func (ll *LinkedList) Sort() {
	if ll.head == nil || ll.head.next == nil {
		return
	}

	var sorted bool
	for !sorted {
		sorted = true
		current := ll.head
		for current.next != nil {
			if fmt.Sprintf("%v", current.data) > fmt.Sprintf("%v", current.next.data) {
				current.data, current.next.data = current.next.data, current.data
				sorted = false
			}
			current = current.next
		}
	}
}

func (ll *LinkedList) Items() []interface{} {
    items := make([]interface{}, 0)
    current := ll.head
    for current != nil {
        items = append(items, current.data)
        current = current.next
    }
    return items
}

func (ll *LinkedList) Contains(data interface{}) bool {
    current := ll.head
    for current != nil {
        if current.data == data {
            return true
        }
        current = current.next
    }
    return false
}

func (ll *LinkedList) Clear() {
    ll.head = nil
}

func (ll *LinkedList) Size() int {
    count := 0
    current := ll.head
    for current != nil {
        count++
        current = current.next
    }
    return count
}
