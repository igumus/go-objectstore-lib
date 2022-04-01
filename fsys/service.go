package fsys

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/igumus/go-objectstore-lib"
)

// Captures/Represents filesystem backed objectstore service information
type fsObjectStoreService struct {
	config *fsObjectStoreConfig
}

// NewFileSystemObjectStore creates file system backed ObjectStore instance via given configuration options.
// Error returns when creating ObjectStore instance failed.
func NewFileSystemObjectStore(opts ...FSObjectstoreConfigOption) (objectstore.ObjectStore, error) {
	cfg := defaultFSObjectstoreConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	if err := cfg.validate(); err != nil {
		return nil, err
	}
	srv := &fsObjectStoreService{
		config: cfg,
	}

	return srv, nil
}

func (a *fsObjectStoreService) checkContextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return objectstore.ErrOperationCancelled
	case context.DeadlineExceeded:
		return objectstore.ErrOperationDeadlineExceeded
	default:
		return nil
	}
}

// toObjectLink - converts cid (aka content identifier) and name to object link for objectstore
func (f *fsObjectStoreService) toObjectLink(cid, name string) string {
	sepLen := f.config.seperatorLength
	parentFolder := string(cid[:sepLen])
	childFolder := string(cid[sepLen:])
	ret := fmt.Sprintf(_objFormat, f.config.dir, f.config.bucket, parentFolder, childFolder, name)
	if f.config.debug {
		log.Printf("debug: objectLink %s, %s: %s\n", cid, name, ret)
	}
	return ret
}

// exists - checks whether specified fileName exists.
func (f *fsObjectStoreService) exists(fileName string) bool {
	_, err := os.Stat(fileName)
	return !errors.Is(err, os.ErrNotExist)
}

// read from fileName
func (f *fsObjectStoreService) read(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("err: opening object failed: %s, %v\n", fileName, err)
		return nil, objectstore.ErrObjectReadingFailed
	}

	binData := bytes.Buffer{}
	_, err = binData.ReadFrom(file)
	if err != nil {
		log.Printf("err: reading object failed: %s, %v\n", fileName, err)
		return nil, objectstore.ErrObjectReadingFailed
	}

	return binData.Bytes(), nil
}

// write - writes data (aka content) to specified fileName
func (f *fsObjectStoreService) write(fileName string, data []byte) error {
	os.MkdirAll(filepath.Dir(fileName), 0777)
	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("err: creating object failed: %s, %v\n", fileName, err)
		return objectstore.ErrObjectWritingFailed
	}

	binData := bytes.Buffer{}
	binData.Write(data)

	_, err = binData.WriteTo(file)
	if err != nil {
		log.Printf("err: writing object failed: %s, %v\n", fileName, err)
		return objectstore.ErrObjectWritingFailed
	}

	return nil
}
