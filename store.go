package objectstore

import "context"

// ObjectStoreReader defines the functions clients need to perform
// reading operations (chunk, block) on objectstore
type ObjectStoreReader interface {
	ReadObject(ctx context.Context, cid string) ([]byte, error)
	HasObject(ctx context.Context, cid string) bool
}

// ObjectStoreWriter defines the functions clients need to perform
// writing operations (object, references) on objectstore
type ObjectStoreWriter interface {
	DigestObject(context.Context, []byte) string
	CreateObject(ctx context.Context, data []byte) (string, error)
}

// Captures/Represents objectstore reference structure information.
type Ref struct {
	Prev string `json:"prev"`
	Curr string `json:"curr"`
}

// ObjectStore defines the functions clients need to perform
// read/write operations on objectstore. It derives reading functions
// from ObjectStoreReader interface, and writing functions from
// ObjectStoreWriter interface
type ObjectStore interface {
	ObjectStoreReader
	ObjectStoreWriter
}
