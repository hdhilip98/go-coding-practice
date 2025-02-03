package ds

import "errors"

type queue[T any] struct {
	data []T
}

func NewQueue[T any]() *queue[T] {
	return &queue[T]{
		data: []T{},
	}
}

func (q *queue[T]) Enqueue(node T) {
	q.data = append(q.data, node)
}

func (q *queue[T]) Dequeue() (T, error) {
	var node T
	if len(q.data) <= 0 {
		return node, errors.New("empty queue")
	}
	node = q.data[0]
	q.data = q.data[1:]
	return node, nil
}

func (q *queue[T]) Peek() (T, error) {
	var node T
	if len(q.data) <= 0 {
		return node, errors.New("empty queue")
	}
	node = q.data[0]
	return node, nil
}
