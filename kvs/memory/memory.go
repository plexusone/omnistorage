// Copyright 2025 John Wang. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package memory re-exports the in-memory key-value storage backend.
//
// This is the default storage backend, suitable for development
// and testing. Data is not persisted across restarts.
package memory

import (
	"github.com/plexusone/omnistorage-core/kvs/backend/memory"
)

// Re-export Store type.
type Store = memory.Store

// New creates a new in-memory storage.
var New = memory.New
