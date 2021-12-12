package main

import (
	"fmt"
)

func main() {
	var num int = 5
	fmt.Println(num, &num)
	squareVal(num)
	fmt.Println(num, &num)
	squareAddr(&num)
	fmt.Println(num, &num)
}

func squareVal(v int) {
	v = v * v
	fmt.Println(v, &v)
}

func squareAddr(p *int) {
	*p = *p * *p
	fmt.Println("--->", p, *p)
}
