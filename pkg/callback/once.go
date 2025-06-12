package callback

import (
	"sync/atomic"
)

var NoError = func(err error) {}

type cb[T any] struct {
	onFinish   func(t T)
	onError    func(err error)
	beforeHook atomic.Pointer[func()]
	finished   atomic.Bool
}

func (c *cb[T]) Finish(t T) {
	if !c.finished.CompareAndSwap(false, true) {
		return
	}
	attachFunc := c.beforeHook.Load()
	if attachFunc != nil {
		(*attachFunc)()
	}
	c.onFinish(t)
}

func (c *cb[T]) BeforeHook(f func()) {
	c.beforeHook.Store(&f)
}

func (c *cb[T]) Error(err error) {
	if !c.finished.CompareAndSwap(false, true) {
		return
	}
	attachFunc := c.beforeHook.Load()
	if attachFunc != nil {
		(*attachFunc)()
	}
	c.onError(err)
}

func NewOnce[T any](onFinish func(t T), onError func(err error)) Callback[T] {
	return &cb[T]{
		onFinish,
		onError,
		atomic.Pointer[func()]{},
		atomic.Bool{},
	}
}
