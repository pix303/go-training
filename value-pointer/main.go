package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

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
	fmt.Printf("foo pointer %p \n ", &f)
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
	fmt.Println("value of pointer", p)
	var v int = 23
	p = &v
	fmt.Println("value of pointer after assignation of value", p, *p, v)
	*p = 45
	fmt.Println("value of pointer and value after setting value by pointer", p, *p, v)

	fmt.Println("")
	fmt.Println("---Declare struct and show same syntax using value or pointer of struct")
	person := Person{"James", 22}
	fmt.Println("Init person", person.Name, person.Age)
	personPointer := &person
	fmt.Printf("Person pointer %p - %v\n", personPointer, personPointer)
	fmt.Println("Init person and see by pointer", *&personPointer.Name, *&personPointer.Age)
	personPointer.Age = 34
	fmt.Println("Change age by pointer without prefix *", person.Name, person.Age)
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
	fmt.Printf("--> value: %v; struct pointer %p \n", f, &f)
}

func (f foo) DecreaseValue() {
	f.counter--
	fmt.Printf("--> value: %v; struct pointer %p \n", f, &f)
}

type CounterPointer interface {
	IncreasePointer()
	DecreasePointer()
}

func (f *foo) IncreasePointer() {
	fmt.Printf("--> struct pointer %p \n", f)
	f.counter++
}

func (f *foo) DecreasePointer() {
	fmt.Printf("--> struct pointer %p \n", f)
	f.counter--
}
