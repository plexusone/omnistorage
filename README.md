# OmniStorage

[![Go Reference](https://pkg.go.dev/badge/github.com/plexusone/omnistorage.svg)](https://pkg.go.dev/github.com/plexusone/omnistorage)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

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

| Backend | Provider | Registry Name |
|---------|----------|---------------|
| Amazon S3 | [omni-aws](https://github.com/plexusone/omni-aws) | `s3` |
| GitHub Releases | [omni-github](https://github.com/plexusone/omni-github) | `github` |
| Google Cloud Storage | [omni-google](https://github.com/plexusone/omni-google) | `gcs` |
| Google Drive | [omni-google](https://github.com/plexusone/omni-google) | `gdrive` |

## Minimal Dependencies

For minimal dependencies, import [omnistorage-core](https://github.com/plexusone/omnistorage-core) directly and only the specific provider backends you need:

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

- [omnistorage-core](https://github.com/plexusone/omnistorage-core) - Core interfaces and types
- [omni-aws](https://github.com/plexusone/omni-aws) - AWS S3 backend
- [omni-github](https://github.com/plexusone/omni-github) - GitHub Releases backend
- [omni-google](https://github.com/plexusone/omni-google) - Google Cloud Storage and Google Drive backends

## License

MIT
