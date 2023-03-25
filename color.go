package logger

import (
	"fmt"
)

type Color uint8

// Foreground colors basic foreground colors 30 - 37
const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta // 品红
	Cyan    // 青色
	White
)

// Extra foreground color 90 - 97(非标准)
const (
	DarkGray Color = iota + 90 // 亮黑（灰）
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightMagenta
	LightCyan
	LightWhite
	Gray = DarkGray // 亮黑（灰）
)

func (c Color) Text(msg string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c, msg)
}
