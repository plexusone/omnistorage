// Copyright 2025 John Wang. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package sqlite re-exports the SQLite key-value storage backend.
//
// This is the recommended storage backend for persistent storage.
// Data is persisted to a local SQLite database file.
package sqlite

import (
	"github.com/plexusone/omnistorage-core/kvs/backend/sqlite"
)

// Re-export types.
type (
	// Store implements kvs.Store with SQLite.
	Store = sqlite.Store

	// Config configures the SQLite storage.
	Config = sqlite.Config
)

// New creates a new SQLite storage.
var New = sqlite.New
