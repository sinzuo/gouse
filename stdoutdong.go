// stdoutdong
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	doneSignal := make(chan bool)
	go func(ch chan bool) {
		progressSigns := []string{"|", "/", "-", "\\", "|"}
		for {
			for _, p := range progressSigns {
				fmt.Print("\rProgress: ", p)
				os.Stdout.Sync()
				select {
				case <-ch:
					return
				case <-time.After(time.Millisecond * 200):
					continue
				}
			}
		}
	}(doneSignal)

	doneSignal <- true
	time.Sleep(time.Second * 100)
}
