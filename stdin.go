package termscroll

import (
	"os"
	"os/exec"
)

func ReadStdin(stdin chan string) {
	// No buffering
	exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()

	// No visible output
	exec.Command("stty", "-f", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		stdin <- string(b)
	}
}
