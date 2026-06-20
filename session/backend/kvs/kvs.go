// Package kvs re-exports the KVS-backed session storage backend.
//
// This backend adapts any kvs.ListableStore (Redis, SQLite, etc.)
// for session storage with user indexing and automatic cleanup.
package kvs

import (
	"github.com/plexusone/omnistorage-core/session/backend/kvs"
)

// Re-export types.
type Store = kvs.Store

// Re-export functions.
var (
	New             = kvs.New
	NewWithControls = kvs.NewWithControls
)
