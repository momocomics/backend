package storage

import (
	"context"
	"io"
)

type Database interface {
	Get(ctx context.Context, path string) (io.ReadCloser, error)
}
