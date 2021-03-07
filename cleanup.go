package termscroll

import (
	"os"
	"os/signal"
	"syscall"
)

func InitExitStrat(doneChan chan bool) (func(), func()) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	listen := func() {
		<-sigs
		doneChan <- true
	}

	cleanup := showCursor

	return listen, cleanup
}
