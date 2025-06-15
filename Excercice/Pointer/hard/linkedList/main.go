package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func printList(head *Node) {
	for head != nil {
		fmt.Printf("%d", head.value)
		if head.next != nil {
			fmt.Print(" -> ")
		}
		head = head.next
	}
	fmt.Println()
}

func deleteNode(head **Node, target int) {
	curr := *head
	var prev *Node

	for curr != nil {
		if curr.value == target {
			if prev == nil {
				*head = curr.next
			} else {
				prev.next = curr.next
			}
			return
		}
		prev = curr
		curr = curr.next
	}
}

func main() {
	n3 := &Node{value: 30}
	n2 := &Node{value: 20, next: n3}
	n1 := &Node{value: 10, next: n2}

	head := n1

	fmt.Print("List: ")
	printList(head)

	fmt.Println("Deleted: 20")
	deleteNode(&head, 20)
	
	deleteNode(&head, 30)
	fmt.Print("List: ")
	printList(head)
}
