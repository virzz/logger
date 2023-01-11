//go:build release || !debug
// +build release !debug

package logger

func SetDebug() {}

func IsDebug() bool {
	return false
}

func Debug(v ...any)                 {}
func DebugF(format string, v ...any) {}
