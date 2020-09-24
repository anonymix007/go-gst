# go-gst

Go bindings for the gstreamer C library

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-rounded)](https://pkg.go.dev/github.com/tinyzimmer/go-gst/gst)
[![godoc reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/tinyzimmer/go-gst/gst)
[![GoReportCard example](https://goreportcard.com/badge/github.com/nanomsg/mangos)](https://goreportcard.com/report/github.com/tinyzimmer/go-gst)

This package was originally written to aid the audio support in [`kvdi`](https://github.com/tinyzimmer/kvdi). 
But it made sense to turn it into an independent, consumable package. The intention now is to progressively implement the entire API.

See the go.dev reference for documentation and examples.

For other examples see the command line implementation [here](cmd/go-gst).

_TODO: Write examples on programatically building the pipeline yourself_

## Requirements

For building applications with this library you need the following:

 - `cgo`: You must set `CGO_ENABLED=1` in your environment when building
 - `pkg-config`
 - `libgstreamer-1.0-dev`: This package name may be different depending on your OS. You need the `gst.h` header files.

For running applications with this library you'll need to have `libgstreamer-1.0` installed. Again, this package may be different depending on your OS.


## CLI

There is a CLI utility included with this package that demonstrates some of the things you can do.

For now the functionality is limitted to GIF encoing and other arbitrary pipelines.
If I extend it further I'll publish releases, but for now, you can retrieve it with `go get`.

```bash
go get github.com/tinyzimmer/go-gst-launch/cmd/go-gst-launch
```

The usage is described below:

```
Go-gst is a CLI utility aiming to implement the core functionality
of the core gstreamer-tools. It's primary purpose is to showcase the functionality of 
the underlying go-gst library.

There are also additional commands showing some of the things you can do with the library,
such as websocket servers reading/writing to/from local audio servers and audio/video/image
encoders/decoders.

Usage:
  go-gst [command]

Available Commands:
  completion  Generate completion script
  gif         Encodes the given video to GIF format
  help        Help about any command
  inspect     Inspect the elements of the given pipeline string
  launch      Run a generic pipeline
  websocket   Run a websocket audio proxy for streaming audio from a pulse server 
              and optionally recording to a virtual mic.

Flags:
  -I, --from-stdin      Write to the pipeline from stdin. If this is specified, then -i is ignored.
  -h, --help            help for go-gst
  -i, --input string    An input file, defaults to the first element in the pipeline.
  -o, --output string   An output file, defaults to the last element in the pipeline.
  -O, --to-stdout       Writes the results from the pipeline to stdout. If this is specified, then -o is ignored.
  -v, --verbose         Verbose output. This is ignored when used with --to-stdout.

Use "go-gst [command] --help" for more information about a command.
```