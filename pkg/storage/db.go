package storage

import (
	"context"
	"io"
)

type DB interface {
	Get(ctx context.Context, path string) (io.ReadCloser, error)
}
