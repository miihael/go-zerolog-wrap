package logwrap

import (
	"bufio"
	"io"
	"log"
	"sync"

	"github.com/rs/zerolog"
)

type LogWrapper struct {
	wg    sync.WaitGroup
	zlg   *zerolog.Logger
	level zerolog.Level
	lg    *log.Logger
	r     *io.PipeReader
	w     *io.PipeWriter
}

func New(zlg *zerolog.Logger, level zerolog.Level) *LogWrapper {
	r, w := io.Pipe()
	scanner := bufio.NewScanner(r)
	lw := &LogWrapper{
		lg:    log.New(w, "", 0),
		zlg:   zlg,
		level: level,
		r:     r,
		w:     w,
	}
	lw.wg.Add(1)
	go func() {
		defer lw.wg.Done()
		for scanner.Scan() {
			lw.zlg.WithLevel(lw.level).Msg(scanner.Text())
		}
	}()
	return lw
}

func (lw *LogWrapper) Logger() *log.Logger {
	return lw.lg
}

func (lw *LogWrapper) Close() {
	lw.r.Close()
	lw.w.Close()
	lw.wg.Wait()
}
