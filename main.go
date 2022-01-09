package main

import (
	"fmt"

	"github.com/shashisharma307/linkedlist/singlylinkedlist"
)

func main() {

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

}
