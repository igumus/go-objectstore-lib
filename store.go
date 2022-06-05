package objectstore

import (
	"context"
	"io"
)

// ObjectStore defines the functions clients need to perform
// read/write/existence/linking operations on objectstore.
type ObjectStore interface {
	CreateObject(context.Context, io.Reader) (string, error)
	ReadObject(context.Context, string) ([]byte, error)
	HasObject(context.Context, string) bool
	ListObject(context.Context) (<-chan string, <-chan error)
}
