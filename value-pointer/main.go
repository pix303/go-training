package main

import "fmt"

func main() {
	// set a variable with name, type and value
	var num int = 5
	fmt.Printf("---Init var with value %v of type %T and pointer address %p---\n", num, num, &num)

	squareValue(num)
	fmt.Println("Calculated square by value of var num", num, &num)

	squarePointer(&num)
	fmt.Println("Calculated square by pointer of var num", num, &num)

	fmt.Println("")
	fmt.Println("---Set value of a struct member by value or pointer")

	var f foo
	f.IncreaseValue()
	fmt.Println("Current couter (value)", f.counter)
	f.DecreaseValue()
	fmt.Println("Current couter (value)", f.counter)
	f.IncreasePointer()
	fmt.Println("Current couter (pointer)", f.counter)
	f.DecreasePointer()
	fmt.Println("Current couter (pointer)", f.counter)

	fmt.Println("")
	fmt.Println("---Declare variable as pointer and var as value")
	var p *int
	fmt.Println("value of pointer",p)
	var v int = 23
	p = &v
	fmt.Println("value of pointer after assignation of value",p,*p,v)
	*p = 45
	fmt.Println("value of pointer and value after setting value by pointer",p,*p,v)


}

// Calc square value of a number
func squareValue(val int) {
	val *= val
	fmt.Println("--> squareValue", val, &val)
}

// Calc square value of a pointer
func squarePointer(val *int) {
	*val *= *val
}

type foo struct {
	counter int
}

type CounterValue interface {
	IncreaseValue()
	DecreaseValue()
}

func (f foo) IncreaseValue() {
	f.counter++
	fmt.Println("--> value and struct pointer", f, &f)
}

func (f foo) DecreaseValue() {
	f.counter--
	fmt.Println("--> value and struct pointer", f, &f)
}

type CounterPointer interface {
	IncreasePointer()
	DecreasePointer()
}

func (f *foo) IncreasePointer() {
	fmt.Println("--> struct pointer", &f)
	f.counter++
}

func (f *foo) DecreasePointer() {
	fmt.Println("--> struct pointer", &f)
	f.counter--
}
