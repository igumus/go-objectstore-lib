package objectstore

import (
	"context"
	"io"
)

// ListObjectEvent - handles async listing objects of objectstore
type ListObjectEvent struct {
	Object string
	Error  error
}

// ObjectStore defines the functions clients need to perform
// read/write/existence/linking operations on objectstore.
type ObjectStore interface {
	CreateObject(context.Context, io.Reader) (string, error)
	ReadObject(context.Context, string) ([]byte, error)
	HasObject(context.Context, string) bool
	ListObject(context.Context) <-chan ListObjectEvent
}
