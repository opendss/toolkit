package callback

import "sync/atomic"

type shardCb[T any] struct {
	onFinish   func(t T)
	onError    func(err error)
	beforeHook atomic.Pointer[func()]
}

func (c *shardCb[T]) Finish(t T) {
	attachFunc := c.beforeHook.Load()
	if attachFunc != nil {
		(*attachFunc)()
	}
	c.onFinish(t)
}

func (c *shardCb[T]) BeforeHook(f func()) {
	c.beforeHook.Store(&f)
}

func (c *shardCb[T]) Error(err error) {
	attachFunc := c.beforeHook.Load()
	if attachFunc != nil {
		(*attachFunc)()
	}
	c.onError(err)
}

func NewShared[T any](onFinish func(t T), onError func(err error)) Callback[T] {
	return &shardCb[T]{
		onFinish,
		onError,
		atomic.Pointer[func()]{},
	}
}
