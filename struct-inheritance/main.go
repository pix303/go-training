package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

func (this *Person) SayHello() {
	fmt.Println("hellloooo im", this.Name)
}

type Player struct {
	Points int
	Person
}

func (this *Player) AddPoints(num int) {
	this.Points += num
	fmt.Println("player points ", this.Name, this.Points)
}

func main() {
	p := Person{
		Name: "pino",
		Age:  33,
	}
	p.SayHello()

	pl := Player{
		Person: Person{
			Name: "franco",
			Age:  33,
		},
		Points: 10,
	}

	pl.SayHello()
	pl.AddPoints(20)

}
