package logger

import (
	"os"
)

type Writer struct {
	f *os.File
}

func (w Writer) Write(p []byte) (n int, err error) {
	if w.f != nil {
		w.f.Write(p)
	}
	return os.Stderr.Write(p)
}
