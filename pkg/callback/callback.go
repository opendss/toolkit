package callback

type Callback[T any] interface {
	Finish(t T)
	Error(err error)
}
