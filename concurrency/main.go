package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println("---------------------------------")
	fmt.Println("------------ in sync ------------")
	FakeTask1()
	FakeTask2()
	FakeTask3()
	FakeTask4()
	fmt.Printf("elapsed time: %v \n", time.Since(now))

	fmt.Println("---------------------------------")
	fmt.Println("---- in async with WaitGroup ----")

	now = time.Now()
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		FakeTask1()
	}()

	go func() {
		defer wg.Done()
		FakeTask2()
	}()

	go func() {
		defer wg.Done()
		FakeTask3()
	}()

	go func() {
		defer wg.Done()
		FakeTask4()
	}()

	wg.Wait()

	fmt.Printf("elapsed time: %v \n", time.Since(now))

	fmt.Println("---------------------------------")
	fmt.Println("---- in async with channels -----")

	now = time.Now()
	var doneChannel = make(chan string)

	go func() {
		FakeTask1()
		doneChannel <- "a"
	}()

	go func() {
		FakeTask2()
		doneChannel <- "b"
	}()

	go func() {
		FakeTask3()
		doneChannel <- "c"
	}()

	go func() {
		FakeTask4()
		doneChannel <- "d"
	}()

	var msg string
	var totalMsg string
	for {
		msg = <-doneChannel
		totalMsg += msg
		if len(totalMsg) == 4 {
			fmt.Println(totalMsg)
			break
		}
	}
	fmt.Printf("elapsed time: %v \n", time.Since(now))

}

// Fake task that take some amount of time
func FakeTask1() {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Task 1 terminated")
}

func FakeTask2() {
	fmt.Println("Task 2 terminated")
}

func FakeTask3() {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Task 3 terminated")
}

func FakeTask4() {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Task 4 terminated")
}
