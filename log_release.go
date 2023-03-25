//go:build !debug
// +build !debug

package logger

func Debug(v ...any)                             {}
func DebugF(format string, v ...any)             {}
func (l *Logger) Debug(v ...any)                 {}
func (l *Logger) DebugF(format string, v ...any) {}
