package objectstore

import "fmt"

const _seperationLength = 2
const _formatLink = "%s/%s/%s"

type LinkFunc func(string) string

var DefaultLinkFunc LinkFunc = func(cid string) string {
	parentFolder := string(cid[:_seperationLength])
	childFolder := string(cid[_seperationLength:])
	return fmt.Sprintf(_formatLink, parentFolder, childFolder, cid)
}
