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
