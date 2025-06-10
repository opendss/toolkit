package io

import (
	"io"

	"go.uber.org/multierr"
)

var _ io.Closer = &closer{}

type closer struct {
	closers []io.Closer
}

func (c *closer) Close() error {
	var err error
	for _, c := range c.closers {
		err = multierr.Append(err, c.Close())
	}
	return err
}

func CloseAll(closers ...io.Closer) io.Closer {
	return &closer{
		closers: closers,
	}
}
