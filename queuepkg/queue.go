package queuepkg

import "fmt"

type data int

type customQueue struct {
	q []data
}

func CreateCustomQueue() *customQueue {
	return &customQueue{
		q: []data{},
	}
}

func (cq *customQueue) Insert(d data) {
	cq.q = append(cq.q, d)
}

func (cq *customQueue) DeQueue() data {
	if len(cq.q) == 0 {
		fmt.Println("Queue is empyt")
	}

	firstItem := cq.q[0]
	cq.q = cq.q[1:]

	return firstItem
}

func (cq *customQueue) Print() {
	if cq.q == nil {
		fmt.Println("Queue is empty")
	}

	for _, v := range cq.q {
		fmt.Print("|", v)
	}
	fmt.Print("|\n")
}
