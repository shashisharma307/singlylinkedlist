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
	fmt.Print(" |")
}
