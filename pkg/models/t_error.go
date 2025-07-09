package models

type TError[T any] struct {
	T   T
	Err error
}
