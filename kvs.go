// Copyright 2025 John Wang. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package omnistorage

import (
	// Re-export core kvs types
	"github.com/plexusone/omnistorage-core/kvs"
)

// Re-export core key-value storage types.
type (
	// Store is the primary key-value storage interface.
	Store = kvs.Store

	// ListableStore extends Store with key listing.
	ListableStore = kvs.ListableStore

	// DocumentStore extends Store with document operations.
	DocumentStore = kvs.DocumentStore

	// Document represents a stored document.
	Document = kvs.Document

	// Config configures storage creation.
	KVSConfig = kvs.Config
)

// Re-export kvs errors.
var (
	ErrKVSNotFound = kvs.ErrNotFound
	ErrKVSClosed   = kvs.ErrClosed
)
