package foo

// Queue is a simple queue implementation
type Queue[T any] struct {
	data []T
	head int
	tail int
}

/*
comment
*/
func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{data: make([]T, size), head: 0, tail: 0}
}
func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
	q.tail++
}
func (q *Queue[T]) Dequeue() T {
	if q.head == q.tail {
		panic("queue is empty")
	}
	v := q.data[q.head]
	q.head++
	if q.head == len(q.data)/2 {
		q.data = q.data[q.head:]
		q.tail -= q.head
		q.head = 0
	}
	return v
}
func (q *Queue[T]) Empty() bool {
	return q.head == q.tail
}
func (q *Queue[T]) Len() int {
	return q.tail - q.head
}
func (q *Queue[T]) Top() T {
	if q.head == q.tail {
		panic("queue is empty")
	}
	return q.data[q.head]
}
