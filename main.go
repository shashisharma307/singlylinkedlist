package main

import (
	"flag"
	"fmt"

	"github.com/shashisharma307/linkedlist/helper"
)

func main() {

	op := flag.String("operation", "NA", "a string")
	flag.Parse()

	switch *op {
	case "singlylinkedlist":
		helper.SinglyLinkedListOperation()
	case "doublylinkedlist":
		helper.DoublyLinkedListOperation()
	case "queue":
		helper.QueueOperation()
	case "circularqueue":
		helper.CircularQueueOperation()
	default:
		fmt.Println("Incorrect operation ")

	}

}
