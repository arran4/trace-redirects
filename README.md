# Trace Redirects

[![Go Report Card](https://goreportcard.com/badge/github.com/arran4/trace-redirects)](https://goreportcard.com/report/github.com/arran4/trace-redirects)
[![Test](https://github.com/arran4/trace-redirects/actions/workflows/test.yml/badge.svg)](https://github.com/arran4/trace-redirects/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/arran4/trace-redirects.svg)](https://pkg.go.dev/github.com/arran4/trace-redirects)

A simple CLI application to trace the redirects on a website. It follows `3xx` redirects and also handles `Refresh` headers (meta-refresh style).

## Features

- Traces standard HTTP 3xx redirects (301, 302, 303, 307, 308).
- Detects `Refresh` headers.
- Handles relative URL redirects.
- Detects redirect loops (stops after 20 redirects).

## Installation

### From Source

Requires Go 1.22+.

```shell
go install github.com/arran4/trace-redirects/cmd/trace-redirects@latest
```

### From Releases

Download the binary for your operating system from the [Releases](https://github.com/arran4/trace-redirects/releases) page.

## Usage

```shell
trace-redirects https://website1.com
```

### Example

```shell
$ trace-redirects https://google.com
2024/05/21 12:00:00 Tracing https://google.com
1: https://google.com
2: https://www.google.com/
Done
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](license.txt)
