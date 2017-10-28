package api

import (
	"io"
	"testing"
)

func TestNewHttpServer(t *testing.T) {
	fn := func(r io.ReadCloser) error {
		return nil
	}

	NewHttpServer(fn)
}
