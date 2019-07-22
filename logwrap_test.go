package logwrap_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/miihael/go-zerolog-wrap"
	"github.com/rs/zerolog"
)

func TestWrap(t *testing.T) {
	var buf bytes.Buffer
	zlg := zerolog.New(&buf).With().Int("test", 123).Logger()
	lw := logwrap.New(&zlg, zerolog.WarnLevel)

	logger := lw.Logger()
	logger.Print("test message")

	lw.Close()

	str := buf.String()
	if !strings.Contains(str, "\"test message\"") || !strings.Contains(str, "123") {
		t.Errorf("Wrapping failed. buffer contents: %s", str)
	} else {
		t.Log(str)
	}
}
