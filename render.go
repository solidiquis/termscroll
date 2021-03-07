package termscroll

import (
	"fmt"
)

func InitRender(sl []string) (func(), func(), func() int) {
	current := 0

	var interval []int
	if len(sl) > 10 {
		interval = []int{0, 10}
	} else {
		interval = []int{0, len(sl)}
	}

	render := func() {
		for i := interval[0]; i < interval[1]; i++ {
			eraseLine()

			if i == current {
				fmt.Println(selected(sl[current]))
				continue
			}

			fmt.Print("  ")
			fmt.Println(sl[i])
		}

		y := 11
		if current < len(sl)-1 {
			fmt.Printf("%s\n", DOWN_POINTER)
		} else {
			y--
			eraseLine()
		}

		cursorUp(y)
	}

	up := func() {
		if current-1 < 0 {
			return
		}

		if current-1 < interval[0] {
			interval[0]--
			interval[1]--
		}
		current--
		render()
	}

	down := func() {
		if current+2 > len(sl) {
			return
		}

		if current+2 > interval[1] {
			interval[0]++
			interval[1]++
		}
		current++
		render()
	}

	getCurrent := func() int {
		return current
	}

	hideCursor()
	render()

	return up, down, getCurrent
}
