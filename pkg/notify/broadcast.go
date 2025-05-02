package notify

import (
	"context"
	"sync"
)

type BroadCast[T any] interface {
	Listen(ctx context.Context) (*T, error)

	BroadCast(t *T)
}

func NewBroadCast[T any]() BroadCast[T] {
	return &broadCast[T]{
		c: make(chan struct{}),
	}
}

type broadCast[T any] struct {
	sync.RWMutex
	t *T
	c chan struct{}
}

func (b *broadCast[T]) Listen(ctx context.Context) (*T, error) {
	b.RLock()
	c := b.c
	b.RUnlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-c:
		b.RLock()
		defer b.RUnlock()
		return b.t, nil
	}
}

func (b *broadCast[T]) BroadCast(t *T) {
	b.Lock()
	b.t = t
	close(b.c)
	b.c = make(chan struct{})
	b.Unlock()
}
