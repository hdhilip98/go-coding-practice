package ds

import "errors"

type stack[T any] struct {
	data []T
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{
		data: []T{},
	}
}

func (s *stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *stack[T]) Pop() (T, error) {
	var item T
	if len(s.data) <= 0 {
		return item, errors.New("empty stack")
	}
	lastIndex := len(s.data) - 1
	item = s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return item, nil
}

func (s *stack[T]) Peek() (T, error) {
	var item T
	if len(s.data) <= 0 {
		return item, errors.New("empty stack")
	}
	item = s.data[len(s.data)-1]
	return item, nil
}
