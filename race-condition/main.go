package main

import (
	"fmt"
	"sync"
	// "time"
)

type Handler struct {
	state  int
	locker sync.Mutex
}

func (h *Handler) ProcessValue(val int) {
	h.locker.Lock()
	defer h.locker.Unlock()

	for i := 0; i < 333455; i++ {
		h.state = val + i
	}
	h.state = val
	fmt.Printf(" val arrived %d\n", h.state)
}

func hello(sendedValue int, wg *sync.WaitGroup) {
	fmt.Printf("sended %d\n", sendedValue)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	hh := Handler{}
	for v := 0; v < 10; v++ {
		wg.Add(1)
		go hh.ProcessValue(v)
		hello(v, &wg)
	}
	wg.Wait()
	// time.Sleep(2 * time.Second)
	fmt.Printf("current state of handler %d\n", hh.state)
}
