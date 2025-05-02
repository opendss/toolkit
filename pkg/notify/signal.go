package notify

import "context"

type Signal interface {
	Listen(ctx context.Context) error

	Signal()
}

var _ Signal = &signal{}

type signal struct {
	c chan struct{}
}

func (s *signal) Listen(ctx context.Context) error {
	select {
	case <-s.c:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *signal) Signal() {
	select {
	case s.c <- struct{}{}:
	default:
	}
}

func NewSignal() Signal {
	return &signal{
		c: make(chan struct{}, 1),
	}
}
