package log

import (
	"bytes"
	"fmt"
	"gf/pkg/log/filewriter"
	"io"
	"path/filepath"
	"runtime"
	"time"
)

const (
	_infoIdx = iota
	_warnIdx
	_errorIdx
	_totalIdx
)

var logTypes = map[int]string{
	_infoIdx:  "info.log",
	_warnIdx:  "warning.log",
	_errorIdx: "error.log",
}

type Config struct {
	Dir	string
}

type LogHandler struct {
	fs    [_totalIdx]*filewriter.FileWriter
}

var handler *LogHandler

func Init(conf *Config) {
	handler = &LogHandler{}
	for idx, file := range logTypes {
		fw, err := filewriter.New(filepath.Join(conf.Dir, file))
		if err != nil {
			panic(err)
		}
		handler.fs[idx] = fw
	}
}

func Close() {
	for idx, _ := range handler.fs {
		handler.fs[idx].Close()
	}
}

func (h *LogHandler) format(text string) string {
	longTime := time.Now().Format("2006-01-02 15:04:05.000")
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "unknow"
	}
	return fmt.Sprintf("[%s] [%s:%d] %s\n", longTime, file, line, text)
}

func (h *LogHandler) log(level int, text string) {
	var (
		w io.Writer
		buf *bytes.Buffer
	)
	buf = &bytes.Buffer{}
	defer func() {
		buf.Reset()
	}()
	w = h.fs[level]
	buf.WriteString(h.format(text))
	buf.WriteTo(w)
}

func Info(format string, args ...interface{}) {
	handler.log(_infoIdx, fmt.Sprintf(format, args...))
}

func Warn(format string, args ...interface{}) {
	handler.log(_warnIdx, fmt.Sprintf(format, args...))
}

func Error(format string, args ...interface{}) {
	handler.log(_errorIdx, fmt.Sprintf(format, args...))
}
