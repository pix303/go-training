package main

import "fmt"

func main() {
	var s = "Ciao 🚨"
	var b = []byte(s)
	var r = []rune(s)

	fmt.Println("----STRING ---------------------------")

	for k, ss := range s {
		fmt.Printf("%d: %q - %x - %U \n", k, ss, ss, ss)
	}

	fmt.Println("--------------------------------------")

	fmt.Println("----BYTES ----------------------------")

	for k, bb := range b {
		fmt.Printf("%d: %q - %x - %U \n", k, bb, bb, bb)
	}

	fmt.Println("--------------------------------------")
	fmt.Println("----RUNES ----------------------------")

	for k, rr := range r {
		fmt.Printf("%d: %q - %x - %U \n", k, rr, rr, rr)
	}

	fmt.Println("--------------------------------------")

}
