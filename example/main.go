package main

import (
	"fmt"
	ts "github.com/solidiquis/termscroll"
)

func main() {
	items := []string{
		"item a",
		"item b",
		"item c",
		"item d",
		"item e",
		"item f",
		"item g",
		"item h",
		"item i",
		"item j",
		"item k",
		"item l",
		"item m",
	}

	done := make(chan bool)
	listenForExitSig, cleanup := ts.InitExitStrat(done)
	go listenForExitSig()

	up, down, current := ts.InitRender(items)

	stdin := make(chan string, 1)
	go ts.ReadStdin(stdin)

readInput:
	for {
		select {
		case <-done:
			break readInput
		case ch := <-stdin:
			switch ch {
			case "k":
				up()
			case "j":
				down()
			case "\n": // Enter
				fmt.Printf("You've selected %s\n", items[current()])
				break readInput
			}
		}
	}

	cleanup()
}
