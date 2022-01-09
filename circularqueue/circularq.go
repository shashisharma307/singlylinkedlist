package circularqueue

import "fmt"

type CircularQueue struct {
	data  []interface{}
	rear  int
	front int
	cap   int
}

func CreateCircularQueue(cap int) *CircularQueue {
	return &CircularQueue{
		data:  make([]interface{}, cap),
		rear:  0,
		front: 0,
		cap:   cap,
	}
}

func (q *CircularQueue) Push(val interface{}) bool {

	if (q.rear+1)%q.cap == q.front {
		return false
	}
	q.data[q.rear] = val
	q.rear = (q.rear + 1) % q.cap
	return true
}

func (q *CircularQueue) Pop() interface{} {
	if q.rear == q.front {
		fmt.Println("Queue is empty")
	}

	data := q.data[q.front]
	q.data[q.front] = nil

	q.front = (q.front + 1) % q.cap
	return data
}

func (q *CircularQueue) QueueLength() int {
	return (q.rear - q.front + q.cap) % q.cap
}

func (q *CircularQueue) CheckIfFull() bool {
	//return (q.rear+1)%q.cap == q.front

	if (q.rear+1)%q.cap == q.front {
		return true
	}
	return false
}

func (q *CircularQueue) CheckIfEmpty() bool {
	//return (q.rear == q.front)

	if q.front == q.rear {
		return true
	}
	return false
}

func (q *CircularQueue) Display() string {
	return fmt.Sprintf("%v", q.data)
}
