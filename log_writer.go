//go:build feat
// +build feat

package logger

import (
	"io"
	"log"
	"os"
)

type LogWriter struct {
	lf *os.File
}

func (lw LogWriter) Write(p []byte) (n int, err error) {
	if lw.lf != nil {
		lw.lf.Write(p)
	}
	return os.Stderr.Write(p)
}

func FileOutput(name string) *log.Logger {
	lf, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		Error(err)
	}
	var lw io.Writer = LogWriter{lf: lf}
	return log.New(lw, "", 0)
}

func LogAsFileOutput(name ...string) {
	_name := "./log.log"
	if len(name) > 0 && len(name[0]) > 0 {
		_name = name[0]
	}
	std = FileOutput(_name)
}
