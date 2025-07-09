package lifecycles

type Cleanup interface {
	Clean() error
}
