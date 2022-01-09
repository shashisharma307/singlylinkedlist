package singlylinkedlist

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedListSingle struct {
	list *Node
}

func CreateNode(val int) *Node {
	return &Node{
		value: val,
		next:  nil,
	}
}

/* 1 - If list is empty than create new node and assign it to head pointer which is LinkedListSingle.list in this case
   2 - If list is not empty then
   	   * create new node
	   * assign next node of newly created node with head which is LinkedListSingle.list in this case
	   * assign head with new node
*/

func (ll *LinkedListSingle) InsertAtBegining(val int) {

	if ll.list == nil {
		ll.list = CreateNode(val)
		return
	}

	temp := CreateNode(val)
	head := ll.list
	temp.next = head
	ll.list = temp
}

func (ll *LinkedListSingle) InsertAtEnd(val int) {
	if ll.list == nil {
		ll.list = CreateNode(val)
		return
	}

	temp := CreateNode(val)
	current := ll.list

	for current.next != nil {
		current = current.next
	}
	current.next = temp
}

func (ll *LinkedListSingle) InsertAtPos(val, pos int) {

	if ll.list == nil || pos == 0 {
		temp := CreateNode(val)
		temp.next = ll.list
		ll.list = temp
		return
	}

	i := 0
	cur := ll.list
	prev := cur

	for cur != nil {
		if i == pos {
			node := CreateNode(val)
			node.next = cur
			prev.next = node
			break
		}

		prev = cur
		cur = cur.next
		i++
	}

}

func (ll *LinkedListSingle) DeleteAtGivenPostion(pos int) {
	if ll.list == nil {
		fmt.Println("List is empty")
		return
	}

	if pos == 1 {
		temp := ll.list
		ll.list = temp.next
		return
	}

	i := 1
	cur := ll.list
	var prev *Node

	for cur != nil {
		if pos == i {
			prev.next = cur.next
			break
		}
		prev = cur
		cur = cur.next
		i++
	}
}

func (ll *LinkedListSingle) DeleteFront() {

	if ll.list == nil {
		fmt.Println("List is empyt")
	}
	cur := ll.list
	ll.list = cur.next
}

func (ll *LinkedListSingle) DeleteEnd() {

	if ll.list == nil {
		fmt.Println("List is empyt")
		return
	}

	temp := ll.list

	if temp.next == nil {
		temp = nil
		ll.list = temp
		return
	}

	for temp.next.next != nil {
		temp = temp.next
	}

	temp.next = nil
}

func (ll *LinkedListSingle) Reverse() *Node {

	head := ll.list

	if head == nil {
		fmt.Println("List is empty")
	}

	var next *Node
	var prev *Node

	current := head

	/*
		Logic to reverse link list
		1. we will maintain 3 pointer := current, prev and next
		2. current will point to current head pointer initially. next and prev initially will be nil
		3. we will loop throught list checking current pointer is not nil
		4. copying current.next to next pointer to maintain the right end of the linkedlist
		5. copying prev pointer to current.next link to establish the current node link with left side of list
		6. copying current pointer value to prev
		7. incrementing current pointer to the next
	*/

	for current != nil {
		next = current.next //Copying the next node value in next pointer
		current.next = prev //In first iteration it will set value nil and from 2nd iteration prev pointer is copied to current.next link
		prev = current      // setting current node to previous so we can have both pointer stored
		current = next      // incrementing the current pointer to next
	}

	return prev
}

func PrintReverse(list *Node) {
	if list == nil {
		fmt.Println("List is empty")
	}

	cur := list
	fmt.Println("**************Reversed LinkedList**************")

	for cur != nil {
		fmt.Printf("| %d ", cur.value)
		cur = cur.next
	}
	fmt.Print(" |\n")
}

func (ll *LinkedListSingle) Print() {

	if ll.list == nil {
		fmt.Println("List is empty")
	}

	cur := ll.list
	fmt.Println("**************LinkedList**************")

	for cur != nil {
		fmt.Printf("| %d ", cur.value)
		cur = cur.next
	}
	fmt.Print(" |\n")
}
