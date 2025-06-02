package lifecycle

type Cleanup interface {
	Clean() error
}
