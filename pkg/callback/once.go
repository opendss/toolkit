package callback

import "sync/atomic"

type cb[T any] struct {
	onFinish func(t T)
	onError  func(err error)
	finished atomic.Bool
}

func (c *cb[T]) Finish(t T) {
	if !c.finished.CompareAndSwap(false, true) {
		return
	}
	c.onFinish(t)
}

func (c *cb[T]) Error(err error) {
	if !c.finished.CompareAndSwap(false, true) {
		return
	}
	c.onError(err)
}

func NewOnce[T any](onFinish func(t T), onError func(err error)) Callback[T] {
	return &cb[T]{
		onFinish,
		onError,
		atomic.Bool{},
	}
}
