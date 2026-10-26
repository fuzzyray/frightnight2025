// Package queue implements a basic queue
package queue

import "fmt"

type Queue struct {
	queue []interface{}
}

// Enqueue an element on to the queue
func (q *Queue) Enqueue(e interface{}) {
	(*q).queue = append((*q).queue, e)
}

// Dequeue the first element from the queue
func (q *Queue) Dequeue() interface{} {
	v := (*q).queue[0]
	(*q).queue = (*q).queue[1:]
	return v
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return len((*q).queue)
}

// IsEmpty returns true if the Queue is empty
func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) String() string {
	return fmt.Sprintf("%v", (*q).queue)
}
