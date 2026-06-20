// Package memory re-exports the in-memory session storage backend.
//
// This backend is suitable for development and testing.
// Data is not persisted across restarts.
package memory

import (
	"github.com/plexusone/omnistorage-core/session/backend/memory"
)

// Re-export types.
type Store = memory.Store

// Re-export functions.
var (
	New             = memory.New
	NewWithControls = memory.NewWithControls
)
