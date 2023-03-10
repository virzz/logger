//go:build debug
// +build debug

package logger

import (
	"log"

	"github.com/gookit/color"
)

var debug = false

func SetDebug(b ...bool) {
	if len(b) > 0 {
		debug = b[0]
	} else {
		debug = !debug
	}
}

func IsDebug() bool {
	return debug
}

func Debug(v ...any) {
	if debug {
		std.SetFlags(log.Lshortfile | log.LstdFlags)
		print(color.LightBlue.Sprint(append([]any{"\nš "}, v...)...))
		std.SetFlags(0)
	}
}
func DebugF(format string, v ...any) {
	if debug {
		std.SetFlags(log.Lshortfile | log.LstdFlags)
		print(color.LightBlue.Sprintf("\nš "+format, v...))
		std.SetFlags(0)
	}
}

// Logger

func (l *Logger) SetDebug(b ...bool) {
	if len(b) > 0 {
		l.debug = b[0]
	} else {
		l.debug = !l.debug
	}
}

func (l *Logger) IsDebug() bool {
	return l.debug
}

func (l *Logger) Debug(v ...any) {
	if l.debug {
		l.std.SetFlags(log.Lshortfile | log.LstdFlags)
		l.print(color.LightBlue.Sprint(append([]any{"\nš "}, v...)...))
		l.std.SetFlags(0)
	}
}
func (l *Logger) DebugF(format string, v ...any) {
	if l.debug {
		l.std.SetFlags(log.Lshortfile | log.LstdFlags)
		l.print(color.LightBlue.Sprintf("\nš "+format, v...))
		l.std.SetFlags(0)
	}
}
