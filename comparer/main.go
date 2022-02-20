package main

import "fmt"

type Hobby struct {
	Name   string
	Rating int
}

type Person struct {
	Name     string
	Age      int
	Lastname string
	Passion  Hobby
}

func main() {
	fmt.Println("Comparer tests")

	var p1 = Person{Name: "Paolo", Age: 49, Lastname: "Carraro", Passion: Hobby{Name: "tennis", Rating: 5}}
	var p2 = Person{Name: "Paolo", Age: 49, Lastname: "Carraro", Passion: Hobby{Name: "tennis", Rating: 5}}

	p3 := new(Person)
	p3.Name = "Paolo"
	p3.Age = 49
	p3.Lastname = "Carraro"
	p3.Passion = Hobby{"tennis", 5}

	p4 := p1
	p5 := &p1

	fmt.Printf("Compare values p1 p2 : %v \n", p1 == p2)
	fmt.Printf("Compare values p1 p3 : %v \n", p1 == *p3)
	fmt.Printf("Compare values p1 p4 : %v \n", p1 == p4)
	fmt.Printf("Compare values p1 p5 : %v \n", p1 == *p5)
	fmt.Printf("Compare pointer p1 p2: %v \n", &p1 == &p2)
	fmt.Printf("Compare pointer p1 p3: %v \n", &p1 == p3)
	fmt.Printf("Compare pointer p1 p4: %v \n", &p1 == &p4)
	fmt.Printf("Compare pointer p1 p5: %v \n", &p1 == p5)

	fmt.Printf("p1 Pointer %p - Value %v \n", &p1, p1)
	fmt.Printf("p2 Pointer %p - Value %v \n", &p2, p2)
	fmt.Printf("p3 Pointer %p - Value %v \n", &p3, *p3)
	fmt.Printf("p4 Pointer %p - Value %v \n", &p4, p4)
	fmt.Printf("p5 Pointer %p - Value %v \n", p5, *p5)

}
