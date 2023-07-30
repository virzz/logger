//go:build debug
// +build debug

package logger

import (
	"log"
)

func (l *Logger) Debug(v ...any) {
	l.SetFlags(log.Llongfile | log.LstdFlags)
	l.print(level_DEBUG, v...)
	l.ResetFlags()
}
func (l *Logger) DebugF(format string, v ...any) {
	l.SetFlags(log.Llongfile | log.LstdFlags)
	l.printf(level_DEBUG, format, v...)
	l.ResetFlags()
}
func Debug(v ...any)                 { std.Debug(v...) }
func DebugF(format string, v ...any) { std.DebugF(format, v...) }
