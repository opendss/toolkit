package callback

import (
	"github.com/opendss/toolkit/pkg/model"
)

type Sync[T any] struct {
	ch chan model.TError[T]
}

func (s *Sync[T]) Finish(t T) {
	s.ch <- model.TError[T]{
		T:   t,
		Err: nil,
	}
}

func (s *Sync[T]) Error(err error) {
	s.ch <- model.TError[T]{
		Err: err,
	}
}

func (s *Sync[T]) Wait() (T, error) {
	t := <-s.ch
	return t.T, t.Err
}

func NewSync[T any]() *Sync[T] {
	return &Sync[T]{
		ch: make(chan model.TError[T], 1),
	}
}
