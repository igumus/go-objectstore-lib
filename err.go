package objectstore

import "errors"

// ErrBucketNotSpecified is return, when bucket not specified.
var ErrBucketNotSpecified = errors.New("objectstore: bucket parameter not specified")

// ErrObjectReadingFailed is return, when reading object failed from objectstore
var ErrObjectReadingFailed = errors.New("objectstore: reading object failed")

// ErrObjectWritingFailed is return, when writing object failed
var ErrObjectWritingFailed = errors.New("objectstore: writing object failed")

// ErrObjectNotExists is return, when object not exists in objectstore
var ErrObjectNotExists = errors.New("objectstore: object not exists")

// ErrReferenceNotExists is return, when reference not exists in objectstore
var ErrReferenceNotExists = errors.New("objectstore: reference not exists")

// ErrReferenceWritingFailed is return, when writing reference failed
var ErrReferenceWritingFailed = errors.New("objectstore: writing reference failed")

// ErrReferenceReadingFailed is return, when reading reference failed
var ErrReferenceReadingFailed = errors.New("objectstore: reading reference failed")

// ErrReferenceDecodingFailed is return, when decoding (json unmarshalling) reference failed
var ErrReferenceDecodingFailed = errors.New("objectstore: decoding reference failed")

// ErrOperationCancelled is return, when objectstore operation is cancelled via callee.
var ErrOperationCancelled = errors.New("objectstore: operation cancelled")

// ErrOperationDeadlineExceeded is return, when objectstore operation deadline is exceed (aka timedout)
var ErrOperationDeadlineExceeded = errors.New("objectstore: operation deadline exceeded")
