# termscroll
<center>
  <img src="https://github.com/solidiquis/termscroll/blob/master/assets/demo.gif">
</center>

## Installation & Usage
`$ go get github.com/solidiquis/termscroll`

```go
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
                // ...
	}

        // Graceful exit necessary to unhide cursor
	done := make(chan bool)
	listenForExitSig, cleanup := ts.InitExitStrat(done)
	go listenForExitSig()

        // up() traverses list up
        // down() traverses list down
        // current() gets current active index
	up, down, current := ts.InitRender(items)

        // Read from stdin without buffering
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
        // unhide cursor 
	cleanup()
}

```
