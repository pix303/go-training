package main

import "fmt"

func main() {
	// set a variable with name, type and value
	var num int = 5
	fmt.Println("Init value and pointer address of num", num, &num)

	squareValue(num)
	fmt.Println("Calculated square by value of num", num, &num)

	squarePointer(&num)
	fmt.Println("Calculated square by pointer of num", num, &num)

}

// Calc square value of a number
func squareValue(val int) {
	val *= val
	fmt.Println("Inside squareValue", val, &val)
}

// Calc square value of a pointer
func squarePointer(val *int) {
	*val *= *val
}
