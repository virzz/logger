package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	level_NORMAL int = iota + 1
	level_INFO
	level_WARN
	level_ERROR
	level_FATAL
	level_SUCCESS
	level_DEBUG
	_prefix int = -1
)

var (
	ico = map[int]string{
		level_NORMAL:  "📃 ",
		level_INFO:    "💡 ",
		level_WARN:    "💊 ",
		level_ERROR:   "❎ ",
		level_FATAL:   "💢 ",
		level_SUCCESS: "✅ ",
		level_DEBUG:   "\n🐛 ",
	}
	c = map[int]color{
		level_NORMAL:  lightWhite,
		level_INFO:    lightCyan,
		level_WARN:    lightYellow,
		level_ERROR:   lightRed,
		level_FATAL:   lightRed,
		level_SUCCESS: lightGreen,
		level_DEBUG:   lightBlue,
		_prefix:       lightMagenta,
	}
)

type Logger struct {
	std         *log.Logger
	depth       int
	currentFlag int
}

var std = NewLogger()

func fileLogger(w io.Writer) *Logger { return &Logger{std: log.New(w, "", 0), depth: 4} }
func GetLogger() *Logger             { return std }
func NewLogger() *Logger             { return fileLogger(os.Stderr) }
func FileLogger(fn ...string) *Logger {
	logFile := "logger.log"
	if len(fn) > 0 && len(fn[0]) > 0 {
		logFile = fn[0]
	}
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	if err != nil {
		return NewLogger()
	}
	return fileLogger(Writer{f: f})
}
func SetFileLogger(fn ...string) { std = FileLogger(fn...) }

func (l *Logger) print(lv int, v ...any) { l.std.Output(l.depth, c[lv].Text(ico[lv]+fmt.Sprint(v...))) }
func (l *Logger) fatal(lv int, v ...any) {
	l.std.Output(l.depth, c[lv].Text(ico[lv]+fmt.Sprint(v...)))
	os.Exit(1)
}

/* ======= Export Logger Function (l *Logger) ======= */

func (l *Logger) SetFileLog(fn string) {
	s := FileLogger(fn)
	l.std = s.std
}

func (l *Logger) SetPrefix(prefix string) {
	if len(prefix) > 0 {
		l.std.SetPrefix(c[_prefix].Text(fmt.Sprintf("[%s] ", prefix)))
	} else {
		l.std.SetPrefix("")
	}
}

func (l *Logger) SetFlags(flag int) {
	l.currentFlag = l.std.Flags()
	l.std.SetFlags(flag)
}
func (l *Logger) GetLogLogger() *log.Logger   { return l.std }
func (l *Logger) SetDevFlags()                { l.SetFlags(log.Llongfile | log.LstdFlags) }
func (l *Logger) SetDefaultFlags()            { l.SetFlags(0) }
func (l *Logger) ResetFlags()                 { l.std.SetFlags(l.currentFlag) }
func (l *Logger) SetOutput(w io.Writer)       { l.std.SetOutput(w) }
func (l *Logger) SetDepth(i int)              { l.depth = i }
func (l *Logger) Normal(v ...any)             { l.print(level_NORMAL, v...) }
func (l *Logger) Info(v ...any)               { l.print(level_INFO, v...) }
func (l *Logger) Warn(v ...any)               { l.print(level_WARN, v...) }
func (l *Logger) Error(v ...any)              { l.print(level_ERROR, v...) }
func (l *Logger) Fatal(v ...any)              { l.fatal(level_FATAL, v...) }
func (l *Logger) Success(v ...any)            { l.print(level_SUCCESS, v...) }
func (l *Logger) NormalF(f string, v ...any)  { l.print(level_NORMAL, fmt.Sprintf(f, v...)) }
func (l *Logger) InfoF(f string, v ...any)    { l.print(level_INFO, fmt.Sprintf(f, v...)) }
func (l *Logger) WarnF(f string, v ...any)    { l.print(level_WARN, fmt.Sprintf(f, v...)) }
func (l *Logger) ErrorF(f string, v ...any)   { l.print(level_ERROR, fmt.Sprintf(f, v...)) }
func (l *Logger) FatalF(f string, v ...any)   { l.fatal(level_FATAL, fmt.Sprintf(f, v...)) }
func (l *Logger) SuccessF(f string, v ...any) { l.print(level_SUCCESS, fmt.Sprintf(f, v...)) }

/* ======= Export Logger Function ======= */

func SetFlags(flag int)           { std.SetFlags(flag) }
func SetDevFlags()                { std.SetDevFlags() }
func ResetFlags()                 { std.ResetFlags() }
func SetOutput(w io.Writer)       { std.SetOutput(w) }
func SetDepth(i int)              { std.SetDepth(i) }
func SetPrefix(prefix string)     { std.SetPrefix(prefix) }
func Normal(v ...any)             { std.Normal(v...) }
func Info(v ...any)               { std.Info(v...) }
func Warn(v ...any)               { std.Warn(v...) }
func Error(v ...any)              { std.Error(v...) }
func Fatal(v ...any)              { std.Fatal(v...) }
func Success(v ...any)            { std.Success(v...) }
func NormalF(f string, v ...any)  { std.NormalF(f, v...) }
func InfoF(f string, v ...any)    { std.InfoF(f, v...) }
func WarnF(f string, v ...any)    { std.WarnF(f, v...) }
func ErrorF(f string, v ...any)   { std.ErrorF(f, v...) }
func FatalF(f string, v ...any)   { std.FatalF(f, v...) }
func SuccessF(f string, v ...any) { std.SuccessF(f, v...) }
