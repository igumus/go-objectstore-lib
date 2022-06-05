package objectstore

import (
	"crypto/sha256"
	"fmt"
)

const _formatHex = "%x"

type DigestFunc func([]byte) string

var DefaultDigestFunc DigestFunc = func(data []byte) string {
	sum := sha256.Sum256(data)
	return fmt.Sprintf(_formatHex, sum)
}
