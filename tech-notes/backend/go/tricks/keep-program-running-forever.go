// Reference: https://stackoverflow.com/a/65127173

package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func forever() {
	for {
		time.Sleep(time.Second)
	}
}

func main() {
	go forever()
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
}
