// Package omnistorage provides a unified storage abstraction layer for Go.
//
// This is an aggregator package that re-exports the core interfaces from
// omnistorage-core and imports all provider backends for automatic registration.
//
// For minimal dependencies, import omnistorage-core directly and only the
// specific provider backends you need.
//
// # Object Storage
//
// For file/blob storage (documents, media, backups):
//
//	import (
//	    "github.com/plexusone/omnistorage"
//	    _ "github.com/plexusone/omni-aws/omnistorage/backend/s3"
//	)
//
//	backend, _ := omnistorage.Open("s3", map[string]string{
//	    "bucket": "my-bucket",
//	    "region": "us-east-1",
//	})
//
// # Key-Value Storage
//
// For session state, caching, and structured data:
//
//	import "github.com/plexusone/omnistorage-core/kvs/backend/sqlite"
//
//	store, _ := sqlite.New(sqlite.Config{Path: "data.db"})
//	store.Set(ctx, "session:123", data, time.Hour)
package omnistorage

import (
	// Re-export core object storage types
	"github.com/plexusone/omnistorage-core/object"

	// Import backends for init() registration
	_ "github.com/plexusone/omni-aws/omnistorage/backend/s3"
	_ "github.com/plexusone/omni-github/backend/github"
	_ "github.com/plexusone/omni-google/omnistorage/backend/drive"
	_ "github.com/plexusone/omni-google/omnistorage/backend/gcs"
)

// Re-export core object storage types.
type (
	// Backend represents a storage backend (S3, GCS, local file, etc.).
	Backend = object.Backend

	// ExtendedBackend extends Backend with additional operations.
	ExtendedBackend = object.ExtendedBackend

	// RecordWriter writes framed records to an underlying writer.
	RecordWriter = object.RecordWriter

	// RecordReader reads framed records from an underlying reader.
	RecordReader = object.RecordReader

	// ObjectInfo contains metadata about a stored object.
	ObjectInfo = object.ObjectInfo

	// Features describes capabilities of a backend.
	Features = object.Features

	// HashType identifies a hash algorithm.
	HashType = object.HashType

	// WriterOption configures a writer.
	WriterOption = object.WriterOption

	// ReaderOption configures a reader.
	ReaderOption = object.ReaderOption
)

// Re-export core functions.
var (
	// Register registers a backend factory.
	Register = object.Register

	// Open creates a backend from the registry.
	Open = object.Open

	// WithContentType sets the content type for a writer.
	WithContentType = object.WithContentType

	// WithMetadata sets metadata for a writer.
	WithMetadata = object.WithMetadata

	// WithOffset sets the read offset.
	WithOffset = object.WithOffset

	// WithLimit sets the read limit.
	WithLimit = object.WithLimit

	// ApplyWriterOptions applies writer options.
	ApplyWriterOptions = object.ApplyWriterOptions

	// ApplyReaderOptions applies reader options.
	ApplyReaderOptions = object.ApplyReaderOptions
)

// Re-export core errors.
var (
	ErrNotFound         = object.ErrNotFound
	ErrPermissionDenied = object.ErrPermissionDenied
	ErrBackendClosed    = object.ErrBackendClosed
	ErrNotSupported     = object.ErrNotSupported
	ErrInvalidPath      = object.ErrInvalidPath
	ErrWriterClosed     = object.ErrWriterClosed
)

// Re-export hash types.
const (
	HashMD5    = object.HashMD5
	HashSHA1   = object.HashSHA1
	HashSHA256 = object.HashSHA256
	HashCRC32C = object.HashCRC32C
)
