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

	fmt.Println("----Insert a string into a string ----------------------------")
	origin := "ABCDEFGHI"
	sb := []byte(origin)
	fmt.Printf("Len and cap of []byte init by string var: %d - %d \n", len(sb), cap(sb))

	sbx := []byte("ABCDEFGHI")
	fmt.Printf("Len and cap of []byte init by literal string : %d - %d \n", len(sbx), cap(sbx))

	toAdd := "x😀z"
	i := 5
	sb = append(sb[:i+len(toAdd)], sb[i:]...)
	fmt.Println(string(sb))
	fmt.Println(len(toAdd))
	for t := i; t < len(toAdd)+i; t++ {
		sb[t] = toAdd[t-i]
	}
	fmt.Println(string(sb))

	fmt.Println("--------------------------------------")
}
