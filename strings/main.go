package main

import "fmt"

func main() {
	var s = "Ciao 🚨 - òàè+"
	var b = []byte(s)
	var r = []rune(s)

	fmt.Println("----STRING for range---------------------------")

	for k, ss := range s {
		fmt.Printf("%d: %q - %x - %U -- %#U \n", k, ss, ss, ss, ss)
	}

	fmt.Println("--------------------------------------")
	fmt.Println("----STRING for loop ----------------------------")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%d: %q - %x - %U -- %#U \n", i, s[i], s[i], s[i], s[i])
	}

	fmt.Println("--------------------------------------")

	fmt.Println("----BYTES ----------------------------")

	for k, bb := range b {
		fmt.Printf("%d: %q - %x - %U \n", k, bb, bb, bb)
	}

	fmt.Println("--------------------------------------")
	fmt.Println("----RUNES for range----------------------------")

	for k, rr := range r {
		fmt.Printf("%d: %q - %x - %U \n", k, rr, rr, rr)
	}

	fmt.Println("--------------------------------------")
	fmt.Println("----RUNES for loop ----------------------------")

	for i := 0; i < len(r); i++ {
		fmt.Printf("%d: %q - %x - %U \n", i, r[i], r[i], r[i])
	}

	fmt.Println("--------------------------------------")

}
