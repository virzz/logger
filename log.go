package logger

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gookit/color"
)

var std = DefaultOutput()

// ======= Config =======

// DefaultOutput New Default Output
func DefaultOutput() *log.Logger {
	return log.New(os.Stderr, "", 0)
}

// SetOutput Set Output
func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

// SetPrefix Set Prefix
func SetPrefix(prefix ...string) {
	if len(prefix) > 0 && prefix[0] != "" {
		std.SetPrefix(color.LightMagenta.Sprintf("[%s] ", prefix[0]))
	} else {
		std.SetPrefix("")
	}
}

// ======= Log to File =======

type LogWriter struct {
	lf *os.File
}

func (lw LogWriter) Write(p []byte) (n int, err error) {
	if lw.lf != nil {
		lw.lf.Write(p)
	}
	return os.Stderr.Write(p)
}

func FileOutput(name string) *log.Logger {
	lf, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		Error(err)
	}
	var lw io.Writer = LogWriter{lf: lf}
	return log.New(lw, "", 0)
}

func LogAsFileOutput(name ...string) {
	_name := "./log.log"
	if len(name) > 0 && len(name[0]) > 0 {
		_name = name[0]
	}
	std = FileOutput(_name)
}

func fatal(v ...any) {
	std.Fatal(v...)
}

func print(v ...any) {
	std.Output(3, fmt.Sprint(v...))
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
	std.Panic(v...)
}
