# OmniStorage

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/plexusone/omnistorage/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/plexusone/omnistorage/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/plexusone/omnistorage/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/plexusone/omnistorage/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/plexusone/omnistorage/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/plexusone/omnistorage/actions/workflows/go-sast-codeql.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/plexusone/omnistorage
 [goreport-url]: https://goreportcard.com/report/github.com/plexusone/omnistorage
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/plexusone/omnistorage
 [docs-godoc-url]: https://pkg.go.dev/github.com/plexusone/omnistorage
 [docs-mkdoc-svg]: https://img.shields.io/badge/Go-dev%20guide-blue.svg
 [docs-mkdoc-url]: https://plexusone.dev/omnistorage
 [viz-svg]: https://img.shields.io/badge/Go-visualizaton-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=plexusone%2Fomnistorage
 [loc-svg]: https://tokei.rs/b1/github/plexusone/omnistorage
 [repo-url]: https://github.com/plexusone/omnistorage
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/plexusone/omnistorage/blob/main/LICENSE

Unified storage abstraction layer for Go. This is an aggregator package that re-exports the core interfaces from [omnistorage-core](https://github.com/plexusone/omnistorage-core) and imports all provider backends for automatic registration.

## Installation

```bash
go get github.com/plexusone/omnistorage
```

## Quick Start

```go
import (
    "github.com/plexusone/omnistorage"
)

// All backends are automatically registered via init()
backend, err := omnistorage.Open("s3", map[string]string{
    "bucket": "my-bucket",
    "region": "us-east-1",
})
```

## Included Backends

All 9 backends are automatically registered when you import this package:

### Core Backends (from omnistorage-core)

| Backend | Registry Name | Description |
|---------|---------------|-------------|
| File | `file` | Local filesystem storage |
| Memory | `memory` | In-memory storage (testing) |
| Channel | `channel` | Go channel-based IPC |
| SFTP | `sftp` | SSH file transfer |
| Dropbox | `dropbox` | Dropbox cloud storage |

### Cloud Backends (from omni-* packages)

| Backend | Provider | Registry Name |
|---------|----------|---------------|
| Amazon S3 | [omni-aws](https://github.com/plexusone/omni-aws) | `s3` |
| GitHub Releases | [omni-github](https://github.com/plexusone/omni-github) | `github` |
| Google Cloud Storage | [omni-google](https://github.com/plexusone/omni-google) | `gcs` |
| Google Drive | [omni-google](https://github.com/plexusone/omni-google) | `drive` |

## Key-Value Storage

KVS interfaces are re-exported for caching, state management, and structured data:

### SQLite Backend

```go
import "github.com/plexusone/omnistorage/kvs/sqlite"

store, _ := sqlite.New(sqlite.Config{Path: "data.db"})
store.Set(ctx, "user:123", userData, 24*time.Hour)
```

### Redis Backend

```go
import "github.com/plexusone/omnistorage/kvs/redis"

store, _ := redis.New(redis.Config{Addr: "localhost:6379"})
store.Set(ctx, "cache:key", data, time.Hour)
```

### In-Memory Backend

```go
import "github.com/plexusone/omnistorage/kvs/memory"

store := memory.New()
store.Set(ctx, "temp:key", data, 5*time.Minute)
```

## Session Storage

Secure, backend-agnostic server-side session management with size limits, JSON validation, and observability hooks:

### In-Memory Sessions (Development)

```go
import (
    "github.com/plexusone/omnistorage/session"
    "github.com/plexusone/omnistorage/session/backend/memory"
)

store := memory.New()
sess := session.NewSession("user-123")
store.Save(ctx, sess)
```

### KVS-Backed Sessions (Production)

Use any KVS backend (Redis, SQLite) for persistent sessions:

```go
import (
    "github.com/plexusone/omnistorage/session"
    "github.com/plexusone/omnistorage/session/backend/kvs"
    "github.com/plexusone/omnistorage/kvs/redis"
)

redisStore, _ := redis.New(redis.DefaultConfig())
sessionStore := kvs.New(redisStore)

sess := session.NewSession("user-123")
sessionStore.Save(ctx, sess)
```

### Session Controls

Add size limits and violation handling:

```go
import "github.com/plexusone/omnistorage/session/backend/memory"

store := memory.NewWithControls(
    session.WithControls(session.Config{
        MaxSessionSize: 1 << 20, // 1MB limit
        MaxSessions:    1000,
    }),
    session.WithViolationHandler(func(e session.ViolationEvent) {
        log.Printf("session violation: %s", e.Type)
    }),
)
```

## Minimal Dependencies

For minimal dependencies, import [omnistorage-core](https://github.com/plexusone/omnistorage-core) directly:

```go
import "github.com/plexusone/omnistorage-core"

// Gets: file, memory, channel, sftp, dropbox (no cloud SDKs)
backend, err := omnistorage.Open("file", map[string]string{"root": "/data"})
```

Or import only specific backends:

```go
import (
    "github.com/plexusone/omnistorage-core/object"
    _ "github.com/plexusone/omni-aws/omnistorage/backend/s3"
)

backend, err := object.Open("s3", map[string]string{
    "bucket": "my-bucket",
    "region": "us-east-1",
})
```

## Documentation

- [omnistorage-core](https://github.com/plexusone/omnistorage-core) - Core interfaces (object, kvs, session)
- [omni-aws](https://github.com/plexusone/omni-aws) - AWS S3 backend
- [omni-github](https://github.com/plexusone/omni-github) - GitHub Releases backend
- [omni-google](https://github.com/plexusone/omni-google) - Google Cloud Storage and Google Drive backends
- [pkg.go.dev](https://pkg.go.dev/github.com/plexusone/omnistorage) - API reference

## License

MIT
