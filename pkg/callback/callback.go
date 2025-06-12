package callback

type Callback[T any] interface {
	BeforeHook(func())

	Finish(t T)

	Error(err error)
}
