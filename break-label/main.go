package main

import (
	"fmt"
)

func create() {
	c := 0
	fmt.Println("loop starting")
runLoop:
	for {
		c++
		fmt.Println("looping ")
		if c > 5 {
			break runLoop
		}
	}
	fmt.Println("loop stopped")

}

func main() {
	fmt.Println("start")
	create()
}
