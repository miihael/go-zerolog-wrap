# go-zerolog-wrap
log.Logger wrapper using zerolog

Example:
```
package main

import (
	"os"

	"github.com/miihael/go-zerolog-wrap"
	"github.com/rs/zerolog"
)

func main() {
	zlg := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	lw := logwrap.New(zlg, zerolog.DebugLevel)
	defer lw.Close()

	logger := lw.Logger()
	logger.Println("test message")
}
```
