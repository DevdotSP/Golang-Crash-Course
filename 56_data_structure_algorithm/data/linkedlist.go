package data

import "fmt"

// Node represents a node in a singly linked list
type Node struct {
	data int    // The value stored in the node
	next *Node  // Pointer to the next node
}

// List represents a singly linked list
type List struct {
	head *Node  // Pointer to the first node in the list
}

// Add appends a new value to the end of the list
func (l *List) Add(value int) {
	newNode := &Node{data: value}

	// If list is empty, new node becomes the head
	if l.head == nil {
		l.head = newNode
		return
	}

	// Traverse to the end of the list
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	// Append the new node
	curr.next = newNode
}

// Remove deletes the first occurrence of a value from the list
func (l *List) Remove(value int) {
	if l.head == nil {
		// List is empty, nothing to remove
		return
	}

	// If head contains the value, move head to next node
	if l.head.data == value {
		l.head = l.head.next
		return
	}

	// Traverse list to find the node before the target node
	curr := l.head
	for curr.next != nil && curr.next.data != value {
		curr = curr.next
	}

	// Remove the node by skipping it
	if curr.next != nil {
		curr.next = curr.next.next
	}
}

// PrintList prints all values in the list
func PrintList(l *List) {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
	}
	fmt.Println()
}
