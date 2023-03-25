package logger

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gookit/color"
)

const (
	NORMAL int = iota + 1
	INFO
	WARN
	ERROR
	FATAL
	SUCCESS
	DEBUG
	PREFIEX int = -1
)

var (
	ICO = map[int]string{
		NORMAL:  "📃 ",
		INFO:    "💡 ",
		WARN:    "💊 ",
		ERROR:   "❎ ",
		FATAL:   "💢 ",
		SUCCESS: "✅ ",
		DEBUG:   "\n🐛 ",
	}
	C = map[int]Color{
		NORMAL:  LightWhite,
		INFO:    LightCyan,
		WARN:    LightYellow,
		ERROR:   LightRed,
		FATAL:   DarkGray,
		SUCCESS: LightGreen,
		DEBUG:   LightBlue,
		PREFIEX: LightMagenta,
	}
)

type Logger struct {
	std         *log.Logger
	depth       int
	currentFlag int
}

var std = NewLogger()

func NewLogger() *Logger                    { return &Logger{std: log.New(os.Stderr, "", 0), depth: 3} }
func GetLogger() *Logger                    { return std }
func (l *Logger) GetLogLogger() *log.Logger { return l.std }

func (l *Logger) fatal(lv int, v ...any) {
	l.std.Output(l.depth, color.Text(ICO[lv]+fmt.Sprint(v...)))
	os.Exit(1)
}
func (l *Logger) print(lv int, v ...any) {
	l.std.Output(l.depth, C[lv].Text(ICO[lv]+fmt.Sprint(v...)))
}

/* ======= Export Logger Function (l *Logger) ======= */

func (l *Logger) SetFlags(flag int) {
	l.currentFlag = l.std.Flags()
	l.std.SetFlags(flag)
}
func (l *Logger) SetDevFlags()          { l.SetFlags(log.Llongfile | log.LstdFlags) }
func (l *Logger) SetDefaultFlags()      { l.SetFlags(0) }
func (l *Logger) ResetFlags()           { l.std.SetFlags(l.currentFlag) }
func (l *Logger) SetOutput(w io.Writer) { l.std.SetOutput(w) }
func (l *Logger) Setdepth(i int)        { l.depth = i }
func (l *Logger) SetPrefix(prefix string) {
	if len(prefix) > 0 {
		l.std.SetPrefix(color.LightMagenta.Sprintf("[%s] ", prefix))
	} else {
		l.std.SetPrefix("")
	}
}

func (l *Logger) Normal(v ...any)             { l.print(NORMAL, v...) }
func (l *Logger) Info(v ...any)               { l.print(INFO, v...) }
func (l *Logger) Warn(v ...any)               { l.print(WARN, v...) }
func (l *Logger) Error(v ...any)              { l.print(ERROR, v...) }
func (l *Logger) Fatal(v ...any)              { l.fatal(FATAL, v...) }
func (l *Logger) Success(v ...any)            { l.print(SUCCESS, v...) }
func (l *Logger) NormalF(f string, v ...any)  { l.print(NORMAL, fmt.Sprintf(f, v...)) }
func (l *Logger) InfoF(f string, v ...any)    { l.print(INFO, fmt.Sprintf(f, v...)) }
func (l *Logger) WarnF(f string, v ...any)    { l.print(WARN, fmt.Sprintf(f, v...)) }
func (l *Logger) ErrorF(f string, v ...any)   { l.print(ERROR, fmt.Sprintf(f, v...)) }
func (l *Logger) FatalF(f string, v ...any)   { l.fatal(FATAL, fmt.Sprintf(f, v...)) }
func (l *Logger) SuccessF(f string, v ...any) { l.print(SUCCESS, fmt.Sprintf(f, v...)) }

/* ======= Export Logger Function ======= */

func SetFlags(flag int)       { std.SetFlags(flag) }
func SetDevFlags()            { std.SetDevFlags() }
func ResetFlags()             { std.ResetFlags() }
func SetOutput(w io.Writer)   { std.SetOutput(w) }
func Setdepth(i int)          { std.Setdepth(i) }
func SetPrefix(prefix string) { std.SetPrefix(prefix) }

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