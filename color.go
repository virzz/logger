package logger

import (
	"fmt"
)

type color uint8

// Foreground colors basic foreground colors 30 - 37
// const (
// 	black color = iota + 30
// 	red
// 	green
// 	yellow
// 	blue
// 	magenta // 品红
// 	cyan    // 青色
// 	white
// )

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
