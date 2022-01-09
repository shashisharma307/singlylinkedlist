package doublylinkedlist

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type LinkedListDoubly struct {
	list *Node
	tail *Node
}

func CreateNode(val int) *Node {
	return &Node{
		data: val,
		prev: nil,
		next: nil,
	}
}

func (ll *LinkedListDoubly) InsertAtBegining(val int) {
	head := ll.list
	temp := CreateNode(val)

	/*If link list is empty then create the node and simply return*/
	if head == nil {
		ll.list = temp
		return
	}

	temp.next = head //setting next pointer on newly created node to the first node
	head.prev = temp //pointing first node's previous link to temp pointer
	ll.list = temp   //setting head to newly created node

	/*Below code is to maintain the tail pointer so we can print the link list in reverse order
	  Here we are maintaing the previous pointer because when cur pointer become the null we can
	  have the last node in previous pointer, so we can set the previous pointer to tail of linkedlist*/

	cur := ll.list //taking list in cur pointer
	var prev *Node //taking previous node
	for cur != nil {
		prev = cur     //copying cur into previous so
		cur = cur.next //incrementing cur pointer to next node
	}

	ll.tail = prev //setting last node to tail of link list which is maintained by the prev node after for loop ends
}

func (ll *LinkedListDoubly) Print() {

	if ll.list == nil {
		fmt.Println("List is empty")
	}

	cur := ll.list
	fmt.Println("**************LinkedList**************")

	for cur != nil {
		fmt.Printf("| %d ", cur.data)
		cur = cur.next
	}
	fmt.Print(" |\n")
}

func (ll *LinkedListDoubly) PrintReverse() {
	if ll.tail == nil {
		fmt.Println("List is empty")
	}

	cur := ll.tail
	fmt.Println("**************LinkedList**************")

	for cur != nil {
		fmt.Printf("| %d ", cur.data)
		cur = cur.prev
	}
	fmt.Print(" |\n")
}
