# vertex-logger (`vlog`)

Vertex Logger (or `vlog`) is a simple logging library for Golang. This library is used in [Vertex](https://github.com/vertex-center/vertex) to log messages to different outputs:

- `stdout`
- `stderr`
- text file
- JSON file

For text and JSON files, the library open a new file every day. This allows to keep a clean history of the logs, without having huge log files for larger projects.

## Installation

```bash
go get github.com/vertex-center/vlog
```

## Usage

```go
package main

import (
	"errors"

	"github.com/vertex-center/vlog"
)

func main() {
	log := vlog.New(
		// Stdout and stderr.
		vlog.WithOutputStd(),
		// Text file.
		vlog.WithOutputFile("logs", vlog.LogFormatText),
		// JSON file.
		vlog.WithOutputFile("logs", vlog.LogFormatJson),
	)
    defer log.Close()

	log.Debug("message", vlog.String("name", "abc"))
	log.Info("message", vlog.String("name", "abc"))
	log.Warn("message", vlog.String("name", "abc"))
	log.Error(errors.New("message"), vlog.String("name", "abc"))
	log.Request("message", vlog.String("name", "abc"))
}
```

## License

[Vertex-Logger](https://github.com/vertex-center/vertex-logger) is released under the [MIT License](./LICENSE.md).
