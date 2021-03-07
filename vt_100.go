package termscroll

import (
	"fmt"
)

const (
	RIGHT_POINTER = "\x1B[36m\u25B6\x1B[0m"
	UP_POINTER    = "\x1B[31m\u25B2\x1B[0m"
	DOWN_POINTER  = "\x1B[31m\u25BC\x1B[0m"
)

func bold(str string) string {
	return fmt.Sprintf("\x1B[1m%s\x1B[0m", str)
}

func cursorUp(n int) {
	fmt.Printf("\x1B[%dA", n)
}

func hideCursor() {
	fmt.Print("\x1B[?25l")
}

func showCursor() {
	fmt.Print("\x1B[?25h")
}

func selected(str string) string {
	return fmt.Sprintf("%s %s", RIGHT_POINTER, bold(str))
}

func eraseLine() {
	fmt.Print("\x1B[2K")
}
