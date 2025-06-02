package main

import "fmt"

func main() {
	a := MainUtilMethod("ok", "ng")

	fmt.Println(a)

	queue := foo_foo_NewQueue[int](0)

	b := &foo_foo_Queue[int]{}

	var c int

	if b.Empty() {
		c = queue.Len()
	}

	fmt.Println(c)

	fmt.Println(foo_glbl_tete_A)

	return
}

// ------------
// merged package: tete, path: foo/glbl/hoge.go
// ------------
var foo_glbl_tete_A int

func init() {
	foo_glbl_tete_A = 0
}

// ------------------
// merged package: main, path: mainutil.go
// -----------------
func MainUtilMethod(a, b string) int {
	if a == "ok" {
		return 0
	} else if b == "no" {
		return 1
	}
	return 2
}

// --------------
// merged package: foo, path: foo/queue.go
// --------------
// Queue is a simple queue implementation
type foo_foo_Queue[T any] struct {
	data []T
	head int
	tail int
}

/*
comment
*/
func foo_foo_NewQueue[T any](size int) *foo_foo_Queue[T] {
	return &foo_foo_Queue[T]{data: make([]T, size), head: 0, tail: 0}
}
func (q *foo_foo_Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
	q.tail++
}
func (q *foo_foo_Queue[T]) Dequeue() T {
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
func (q *foo_foo_Queue[T]) Empty() bool {
	return q.head == q.tail
}
func (q *foo_foo_Queue[T]) Len() int {
	return q.tail - q.head
}
func (q *foo_foo_Queue[T]) Top() T {
	if q.head == q.tail {
		panic("queue is empty")
	}
	return q.data[q.head]
}
