package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

const (
	timeout = 5000
)

// Learning context by example: slow process that can timeouted/cancelled by context
func main() {

	// rootCtx is the main context created and parent of new context that serves the purpose of timeouting
	// rootCtx can't be cancelled
	rootCtx := context.Background()
	timeoutCtx, cancelFunc := context.WithTimeout(rootCtx, time.Duration(timeout)*time.Millisecond)
	defer cancelFunc()

	externalCancel := make(chan os.Signal)
	signal.Notify(externalCancel, os.Interrupt)

	reciver := make(chan string)

	for i := 0; i < 5; i++ {
		go slowProcessAsync(reciver)
	}

	counter := 0
	for {
		select {
		case res := <-reciver:
			fmt.Println(res)
			counter++
			if counter == 4 {
				close(reciver)
				os.Exit(0)
			}
		case <-timeoutCtx.Done():
			fmt.Println("timeout!")
			close(reciver)
			os.Exit(1)

		case <-externalCancel:
			fmt.Println("stopped by user!")
			close(externalCancel)
			os.Exit(0)
		}
	}

}

func slowProcessAsync(c chan string) {
	c <- slowProcess()
}

func slowProcess() string {
	rand.Seed(time.Now().UnixNano())
	latency := rand.Intn(10000-50) + 50
	time.Sleep(time.Duration(latency) * time.Millisecond)
	return fmt.Sprintf("done after %dms!", latency)
}
