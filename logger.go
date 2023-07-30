package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	LoggerPrint = 1 << 0
	LoggerFile  = 1 << 1
)

type Logger struct {
	depth       int
	currentFlag int
	pLog        *log.Logger
	fLog        *log.Logger
}

var std = NewLogger(LoggerPrint)

func GetLogger() *Logger { return std }
func Default() *Logger {
	return NewLogger(LoggerPrint)
}
func NewLogger(typ int, fn ...string) *Logger {
	l := &Logger{depth: 4}
	switch typ {
	case LoggerPrint:
		l.pLog = log.New(os.Stderr, "", 0)
	case LoggerFile:
		l.fLog = loggerFile(fn...)
	case LoggerPrint | LoggerFile:
		l.pLog = log.New(os.Stderr, "", 0)
		l.fLog = loggerFile(fn...)
	default:
		l.pLog = log.New(os.Stderr, "", 0)
	}
	return l
}

func loggerFile(fn ...string) *log.Logger {
	logFile := "logger.log"
	if len(fn) > 0 && len(fn[0]) > 0 {
		logFile = fn[0]
	}
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	if err != nil {
		Default().Error(err)
		return nil
	}
	return log.New(f, "", 0)
}

func (l *Logger) print(lv int, v ...any) {
	var output = ico[lv] + fmt.Sprint(v...)
	if l.pLog != nil {
		l.pLog.Output(l.depth, c[lv].Text(output))
	}
	if l.fLog != nil {
		l.fLog.Output(l.depth, output)
	}
}

func (l *Logger) printf(lv int, f string, v ...any) {
	var output = ico[lv] + fmt.Sprintf(f, v...)
	if l.pLog != nil {
		l.pLog.Output(l.depth, c[lv].Text(output))
	}
	if l.fLog != nil {
		l.fLog.Output(l.depth, output)
	}
}

func (l *Logger) fatal(lv int, v ...any) {
	l.print(lv, v...)
	os.Exit(1)
}
func (l *Logger) fatalf(lv int, f string, v ...any) {
	l.printf(lv, f, v...)
	os.Exit(1)
}

/* ======= Export Logger Function (l *Logger) ======= */

func (l *Logger) SetPrefix(prefix string) {
	if len(prefix) > 0 {
		if l.pLog != nil {
			l.pLog.SetPrefix(c[_prefix].Text("[" + prefix + "] "))
		}
		if l.fLog != nil {
			l.fLog.SetPrefix("[" + prefix + "] ")
		}
	} else {
		if l.pLog != nil {
			l.pLog.SetPrefix("")
		}
		if l.fLog != nil {
			l.fLog.SetPrefix("")
		}
	}
}

func (l *Logger) SetFlags(flag int) {
	if l.pLog != nil {
		l.currentFlag = l.pLog.Flags()
		l.pLog.SetFlags(flag)
	}
	if l.fLog != nil {
		l.currentFlag = l.fLog.Flags()
		l.fLog.SetFlags(flag)
	}
}

func (l *Logger) SetFile(filename string) {
	l.fLog = loggerFile(filename)
}
func (l *Logger) NoFile() {
	l.fLog.Writer().(*os.File).Close()
	l.fLog = nil
}

func (l *Logger) SetDevFlags()                { l.SetFlags(log.Llongfile | log.LstdFlags) }
func (l *Logger) SetDefaultFlags()            { l.SetFlags(0) }
func (l *Logger) ResetFlags()                 { l.SetFlags(l.currentFlag) }
func (l *Logger) SetDepth(i int)              { l.depth = i }
func (l *Logger) Normal(v ...any)             { l.print(level_NORMAL, v...) }
func (l *Logger) Info(v ...any)               { l.print(level_INFO, v...) }
func (l *Logger) Warn(v ...any)               { l.print(level_WARN, v...) }
func (l *Logger) Error(v ...any)              { l.print(level_ERROR, v...) }
func (l *Logger) Fatal(v ...any)              { l.fatal(level_FATAL, v...) }
func (l *Logger) Success(v ...any)            { l.print(level_SUCCESS, v...) }
func (l *Logger) NormalF(f string, v ...any)  { l.printf(level_NORMAL, f, v...) }
func (l *Logger) InfoF(f string, v ...any)    { l.printf(level_INFO, f, v...) }
func (l *Logger) WarnF(f string, v ...any)    { l.printf(level_WARN, f, v...) }
func (l *Logger) ErrorF(f string, v ...any)   { l.printf(level_ERROR, f, v...) }
func (l *Logger) FatalF(f string, v ...any)   { l.fatalf(level_FATAL, f, v...) }
func (l *Logger) SuccessF(f string, v ...any) { l.printf(level_SUCCESS, f, v...) }

/* ======= Export Logger Function ======= */

func SetFlags(flag int)           { std.SetFlags(flag) }
func SetDevFlags()                { std.SetDevFlags() }
func ResetFlags()                 { std.ResetFlags() }
func SetDepth(i int)              { std.SetDepth(i) }
func SetPrefix(prefix string)     { std.SetPrefix(prefix) }
func SetFile(filename string)     { std.SetFile(filename) }
func NoFile()                     { std.NoFile() }
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
