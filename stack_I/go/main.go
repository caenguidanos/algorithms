package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"sync"
)

type stack[V any] struct {
	mutex sync.Mutex
	data  []V
}

type Stack[V any] interface {
	Pop() error
	Push(value V)
	Top() (V, error)
	Size() int
	IsEmpty() bool
	Memory() int
}

func NewStack[V any]() Stack[V] {
	return &stack[V]{}
}

func (s *stack[V]) Pop() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.IsEmpty() {
		return errors.New("stack is empty")
	}

	var zero V

	hi := len(s.data) - 1
	s.data[hi] = zero
	s.data = s.data[:hi]

	return nil
}

func (s *stack[V]) Push(value V) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.data = append(s.data, value)
}

func (s *stack[V]) Top() (V, error) {
	if s.IsEmpty() {
		var zero V
		return zero, errors.New("stack is empty")
	}

	hi := len(s.data) - 1

	return s.data[hi], nil
}

func (s *stack[V]) Size() int {
	return len(s.data)
}

func (s *stack[V]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *stack[V]) Memory() int {
	b := new(bytes.Buffer)

	if err := gob.NewEncoder(b).Encode(s.data); err != nil {
		panic(err)
	}

	return b.Len()
}
