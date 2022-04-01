package fsys

import (
	"context"
	"log"

	"github.com/igumus/go-objectstore-lib"
)

// HasObject - checks whether object exists on file system with specified cid (aka content identifier)
func (f *fsObjectStoreService) HasObject(ctx context.Context, cid string) bool {
	objLink := f.toObjectLink(cid, cid)
	ret := f.exists(objLink)
	if f.config.debug {
		log.Printf("debug: has object: %s, %t\n", objLink, ret)
	}
	return ret
}

// ReadObject - reads object on file system with specified cid (aka content identifier)
func (f *fsObjectStoreService) ReadObject(ctx context.Context, cid string) ([]byte, error) {
	if !f.HasObject(ctx, cid) {
		return nil, objectstore.ErrObjectNotExists
	}
	objLink := f.toObjectLink(cid, cid)
	if f.config.debug {
		log.Printf("debug: check object existence: %s\n", objLink)
	}
	if ctxErr := f.checkContextError(ctx); ctxErr != nil {
		return nil, ctxErr
	}
	if f.config.debug {
		log.Printf("debug: check context error: %s\n", objLink)
	}
	return f.read(objLink)
}
