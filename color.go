package logger

import (
	"fmt"
)

type color uint8

// Extra foreground color 90 - 97(非标准)
const (
	darkGray color = iota + 90 // 亮黑（灰）
	lightRed
	lightGreen
	lightYellow
	lightBlue
	lightMagenta
	lightCyan
	lightWhite
	gray = darkGray // 亮黑（灰）
)

func (c color) Text(msg string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c, msg)
}

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
