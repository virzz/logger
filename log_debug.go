//go:build debug
// +build debug

package logger

import (
	"fmt"
	"log"
)

func (l *Logger) Debug(v ...any) {
	l.SetFlags(log.Llongfile | log.LstdFlags)
	l.print(L_DEBUG, v...)
	l.ResetFlags()
}
func (l *Logger) DebugF(format string, v ...any) {
	l.SetFlags(log.Llongfile | log.LstdFlags)
	l.print(L_DEBUG, fmt.Sprintf(format, v...))
	l.ResetFlags()
}
func Debug(v ...any)                 { std.Debug(v...) }
func DebugF(format string, v ...any) { std.DebugF(format, v...) }
