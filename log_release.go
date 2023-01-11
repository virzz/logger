//go:build !debug
// +build !debug

package logger

func SetDebug(b ...bool) {}

func IsDebug() bool {
	return false
}

func Debug(v ...any)                 {}
func DebugF(format string, v ...any) {}
