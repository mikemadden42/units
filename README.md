# units

![example workflow](https://github.com/mikemadden42/units/actions/workflows/go.yml/badge.svg)

A small CLI that converts a byte count into a human-readable size.

## Install

```sh
go install github.com/mikemadden42/units@latest
```

Requires Go 1.26 or newer.

## Usage

```sh
$ units -size 1048576
1M
$ units -size 1536
1.5K
```

## Development

```sh
make check    # golangci-lint
make build    # goreleaser snapshot build
go test ./...
```
