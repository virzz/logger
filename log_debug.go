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
		print(color.LightBlue.Sprint(append([]any{"\n🐛 "}, v...)...))
		std.SetFlags(0)
	}
}
func DebugF(format string, v ...any) {
	if debug {
		std.SetFlags(log.Lshortfile | log.LstdFlags)
		print(color.LightBlue.Sprintf("\n🐛 "+format, v...))
		std.SetFlags(0)
	}
}
