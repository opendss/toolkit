package model

type TError[T any] struct {
	T   any
	err error
}
