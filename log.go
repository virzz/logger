//go:build !v2
// +build !v2

package logger

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gookit/color"
)

var calldepth = 3

var std = DefaultOutput()

// ======= Config =======

// DefaultOutput New Default Output
func DefaultOutput() *log.Logger {
	return log.New(os.Stderr, "", 0)
}

// SetOutput Set Output
func SetFlags(flag int) {
	std.SetFlags(flag)
}

// SetOutput Set Output
func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

// SetCalldepth Set Calldepth
func SetCalldepth(i int) {
	calldepth = i
}

// SetPrefix Set Prefix
func SetPrefix(prefix ...string) {
	if len(prefix) > 0 && prefix[0] != "" {
		std.SetPrefix(color.LightMagenta.Sprintf("[%s] ", prefix[0]))
	} else {
		std.SetPrefix("")
	}
}

func fatal(v ...any) {
	std.Output(calldepth, fmt.Sprint(v...))
	os.Exit(1)
}

func print(v ...any) {
	std.Output(calldepth, fmt.Sprint(v...))
}

// ======= Export Logger Function =======

func Normal(v ...any) {
	print(color.LightWhite.Sprint(append([]any{"📃 "}, v...)...))
}
func NormalF(format string, v ...any) {
	print(color.LightWhite.Sprintf("📃 "+format, v...))
}

func Info(v ...any) {
	print(color.LightCyan.Sprint(append([]any{"💡 "}, v...)...))
}
func InfoF(format string, v ...any) {
	print(color.LightCyan.Sprintf("💡 "+format, v...))
}

func Warn(v ...any) {
	print(color.LightYellow.Sprint(append([]any{"💊 "}, v...)...))
}
func WarnF(format string, v ...any) {
	print(color.LightYellow.Sprintf("💊 "+format, v...))
}

func Success(v ...any) {
	print(color.LightGreen.Sprint(append([]any{"✅ "}, v...)...))
}
func SuccessF(format string, v ...any) {
	print(color.LightGreen.Sprintf("✅ "+format, v...))
}

func Error(v ...any) {
	print(color.LightRed.Sprint(append([]any{"❎ "}, v...)...))
}
func ErrorF(format string, v ...any) {
	print(color.LightRed.Sprintf("❎ "+format, v...))
}

func Fatal(v ...any) {
	fatal(color.LightMagenta.Sprint(append([]any{"💢 "}, v...)...))
}
func FatalF(format string, v ...any) {
	fatal(color.LightMagenta.Sprintf("💢 "+format, v...))
}

func Panic(v ...any) {
	std.Output(calldepth, fmt.Sprint(v...))
	panic(v)
}
