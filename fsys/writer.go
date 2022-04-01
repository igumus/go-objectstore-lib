package fsys

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
)

// DigestObject - calculates SHA256 sum of given data.
func (f *fsObjectStoreService) DigestObject(ctx context.Context, data []byte) string {
	sum := sha256.Sum256(data)
	return fmt.Sprintf("%x", sum)
}

// CreateObject - creates object to file system with specified data (aka content)
func (f *fsObjectStoreService) CreateObject(ctx context.Context, data []byte) (string, error) {
	cid := f.DigestObject(ctx, data)
	if err := f.checkContextError(ctx); err != nil {
		return cid, err
	}
	if f.config.debug {
		log.Printf("debug: created object cid: %s\n", cid)
		log.Printf("debug: checked context error: %s\n", cid)
	}
	if !f.HasObject(ctx, cid) {
		objLink := f.toObjectLink(cid, cid)
		return cid, f.write(objLink, data)
	} else if f.config.debug {
		log.Printf("debug: skip writing already exists object: %s\n", cid)
	}
	return cid, nil
}
