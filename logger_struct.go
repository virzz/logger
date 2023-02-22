package logger

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gookit/color"
)

type Logger struct {
	std       *log.Logger
	calldepth int
	//lint:ignore U1000 use in debug
	debug bool
}

var stdV2 = NewLogger()

func NewLogger() *Logger {
	return &Logger{
		std:       log.New(os.Stderr, "", 0),
		calldepth: 3,
	}
}

func GetLogger() *Logger { return stdV2 }

func (l *Logger) GetLogger() *log.Logger {
	return l.std
}

// SetOutput Set Output
func (l *Logger) SetFlags(flag int) {
	l.std.SetFlags(flag)
}

// SetOutput Set Output
func (l *Logger) SetOutput(w io.Writer) {
	l.std.SetOutput(w)
}

// SetCalldepth Set Calldepth
func (l *Logger) SetCalldepth(i int) {
	l.calldepth = i
}

// SetPrefix Set Prefix
func (l *Logger) SetPrefix(prefix ...string) {
	if len(prefix) > 0 && prefix[0] != "" {
		l.std.SetPrefix(color.LightMagenta.Sprintf("[%s] ", prefix[0]))
	} else {
		l.std.SetPrefix("")
	}
}

func (l *Logger) fatal(v ...any) {
	l.std.Output(l.calldepth, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Logger) print(v ...any) {
	l.std.Output(l.calldepth, fmt.Sprint(v...))
}

// ======= Export Logger Func (l *Logger) tion =======

func (l *Logger) Normal(v ...any) {
	l.print(color.LightWhite.Sprint(append([]any{"📃 "}, v...)...))
}
func (l *Logger) NormalF(format string, v ...any) {
	l.print(color.LightWhite.Sprintf("📃 "+format, v...))
}

func (l *Logger) Info(v ...any) {
	l.print(color.LightCyan.Sprint(append([]any{"💡 "}, v...)...))
}
func (l *Logger) InfoF(format string, v ...any) {
	l.print(color.LightCyan.Sprintf("💡 "+format, v...))
}

func (l *Logger) Warn(v ...any) {
	l.print(color.LightYellow.Sprint(append([]any{"💊 "}, v...)...))
}
func (l *Logger) WarnF(format string, v ...any) {
	l.print(color.LightYellow.Sprintf("💊 "+format, v...))
}

func (l *Logger) Success(v ...any) {
	l.print(color.LightGreen.Sprint(append([]any{"✅ "}, v...)...))
}
func (l *Logger) SuccessF(format string, v ...any) {
	l.print(color.LightGreen.Sprintf("✅ "+format, v...))
}

func (l *Logger) Error(v ...any) {
	l.print(color.LightRed.Sprint(append([]any{"❎ "}, v...)...))
}
func (l *Logger) ErrorF(format string, v ...any) {
	l.print(color.LightRed.Sprintf("❎ "+format, v...))
}

func (l *Logger) Fatal(v ...any) {
	l.fatal(color.LightMagenta.Sprint(append([]any{"💢 "}, v...)...))
}
func (l *Logger) FatalF(format string, v ...any) {
	l.fatal(color.LightMagenta.Sprintf("💢 "+format, v...))
}

func (l *Logger) Panic(v ...any) {
	l.std.Output(calldepth, fmt.Sprint(v...))
	panic(v)
}
