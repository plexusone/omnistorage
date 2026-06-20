// Package redis re-exports the Redis key-value storage backend.
//
// This backend is suitable for production deployments requiring
// persistence and multi-server support.
package redis

import (
	"github.com/plexusone/omnistorage-core/kvs/backend/redis"
)

// Re-export types.
type (
	Store  = redis.Store
	Config = redis.Config
)

// Re-export functions.
var (
	New           = redis.New
	DefaultConfig = redis.DefaultConfig
)
