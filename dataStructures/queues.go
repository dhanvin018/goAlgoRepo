package dataStructures

import fmt "fmt"

type Queue struct {
	items    []int
	front    int
	rear     int
	capacity int
}

func NewQueue(capacity int) *Queue {
	if capacity < 1 {
		return nil
	}
	return &Queue{
		items:    make([]int, capacity),
		front:    0,
		rear:     0,
		capacity: capacity,
	}
}

func (q *Queue) Enqueue(value int) error {
	if (q.rear+1)%q.capacity == q.front {
		return fmt.Errorf("queue is full")
	}
	q.items[q.rear] = value
	q.rear = (q.rear + 1) % q.capacity
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.rear == q.front {
		return 0, fmt.Errorf("queue is empty")
	}
	value := q.items[q.front]
	q.front = (q.front + 1) % q.capacity

	return value, nil
}
