//go:build !debug
// +build !debug

package logger

func SetDebug(b ...bool)             {}
func IsDebug() bool                  { return false }
func Debug(v ...any)                 {}
func DebugF(format string, v ...any) {}

// Logger

func (l *Logger) SetDebug(b ...bool)             {}
func (l *Logger) IsDebug() bool                  { return false }
func (l *Logger) Debug(v ...any)                 {}
func (l *Logger) DebugF(format string, v ...any) {}
