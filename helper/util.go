package helper

import (
	"fmt"

	"github.com/shashisharma307/linkedlist/circularqueue"
	"github.com/shashisharma307/linkedlist/doublylinkedlist"
	"github.com/shashisharma307/linkedlist/queuepkg"
	"github.com/shashisharma307/linkedlist/singlylinkedlist"
)

func SinglyLinkedListOperation() {
	ll := singlylinkedlist.LinkedListSingle{}

	fmt.Println("Insert Values")
	ll.InsertAtBegining(56)
	ll.InsertAtBegining(76)
	ll.InsertAtBegining(43)
	ll.InsertAtBegining(46)
	ll.InsertAtBegining(58)

	ll.InsertAtEnd(1)
	ll.InsertAtEnd(2)
	ll.InsertAtEnd(3)
	ll.InsertAtEnd(4)
	ll.InsertAtEnd(5)

	ll.Print()

	ll.InsertAtPos(89, 0)
	ll.Print()

	ll.InsertAtPos(72, 1)
	ll.Print()

	ll.InsertAtPos(74, 2)
	ll.Print()

	ll.DeleteFront()
	ll.Print()

	ll.DeleteEnd()
	ll.Print()

	ll.DeleteAtGivenPostion(3)
	ll.Print()

	ll.DeleteAtGivenPostion(1)
	ll.Print()

	ll.DeleteAtGivenPostion(9)
	ll.Print()

	singlylinkedlist.PrintReverse(ll.Reverse())

}

func DoublyLinkedListOperation() {

	dll := doublylinkedlist.LinkedListDoubly{}
	dll.InsertAtBegining(56)
	dll.InsertAtBegining(76)
	dll.InsertAtBegining(43)
	dll.InsertAtBegining(46)
	dll.InsertAtBegining(58)

	dll.Print()
	dll.PrintReverse()

	dll.InsertAtEnd(45)
	dll.InsertAtEnd(34)
	dll.InsertAtEnd(23)
	dll.InsertAtEnd(54)
	dll.InsertAtEnd(87)

	dll.Print()
	dll.PrintReverse()
}

func QueueOperation() {
	cq := queuepkg.CreateCustomQueue()
	cq.Insert(1)
	cq.Insert(2)
	cq.Insert(3)
	cq.Insert(4)
	cq.Insert(5)

	cq.Print()

	fmt.Println("DeQueued Value :", cq.DeQueue())
	cq.Print()

}

func CircularQueueOperation() {
	c := circularqueue.CreateCircularQueue(3)
	fmt.Println(c.Push(1))
	fmt.Println(c.Push(2))
	fmt.Println(c.Push(3))
	fmt.Println(c.Push(4))
	fmt.Println("realLength", c.QueueLength(), "queue", c.Display())
	fmt.Println(c.Pop())
	fmt.Println("realLength", c.QueueLength(), "queue", c.Display())
	c.Push(9)
	fmt.Println("realLength", c.QueueLength(), "queue", c.Display())
	fmt.Println(c.Pop())
	fmt.Println("realLength", c.QueueLength(), "queue", c.Display())
	c.Push(11)
	fmt.Println("realLength", c.QueueLength(), "queue", c.Display())
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())
	fmt.Println("realLength", c.QueueLength(), "queue", c.Display())
	fmt.Println(c.CheckIfFull(), c.CheckIfEmpty())

}
