package callbacks

import (
	"github.com/opendss/toolkit/pkg/models"
)

type Sync[T any] struct {
	ch chan models.TError[T]
}

func (s *Sync[T]) Finish(t T) {
	s.ch <- models.TError[T]{
		T:   t,
		Err: nil,
	}
}

func (s *Sync[T]) BeforeHook(func()) {
}

func (s *Sync[T]) Error(err error) {
	s.ch <- models.TError[T]{
		Err: err,
	}
}

func (s *Sync[T]) Wait() (T, error) {
	t := <-s.ch
	return t.T, t.Err
}

func NewSync[T any]() *Sync[T] {
	return &Sync[T]{
		ch: make(chan models.TError[T], 1),
	}
}
