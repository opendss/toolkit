package model

type TError[T any] struct {
	T   any
	Err error
}
