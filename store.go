//go:generate mockgen -source store.go -destination mock/store_mock.go -package mock
package objectstore

import (
	"context"
	"io"

	"github.com/ipfs/go-cid"
)

// ListObjectEvent - handles async listing objects of objectstore
type ListObjectEvent struct {
	Object string
	Error  error
}

// ObjectStore defines the functions clients need to perform
// read/write/existence/linking operations on objectstore.
type ObjectStore interface {
	CreateObject(context.Context, io.Reader) (cid.Cid, error)
	ReadObject(context.Context, cid.Cid) ([]byte, error)
	HasObject(context.Context, cid.Cid) bool
	ListObject(context.Context) <-chan ListObjectEvent
}
