package objectstore

import "context"

// ObjectStore defines the functions clients need to perform
// read/write/existence/linking operations on objectstore.
type ObjectStore interface {
	ToObjectLink(cid, name string) string
	DigestObject(context.Context, []byte) string
	CreateObject(ctx context.Context, data []byte) (string, error)
	ReadObject(ctx context.Context, cid string) ([]byte, error)
	HasObject(ctx context.Context, cid string) bool
}
