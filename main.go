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

	fmt.Println("Printing")
	ll.Print()

	ll.InsertAtPos(89, 0)
	fmt.Println("Printing")
	ll.Print()

	ll.InsertAtPos(72, 1)
	fmt.Println("Printing")
	ll.Print()

	ll.InsertAtPos(74, 2)
	fmt.Println("Printing")
	ll.Print()

}
