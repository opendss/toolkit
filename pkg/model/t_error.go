package model

type TError[T any] struct {
	T   T
	Err error
}
