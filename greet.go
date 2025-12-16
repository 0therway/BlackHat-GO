package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) SayHello() {
	fmt.Println("Hello my name is,", p.Name)

}

type Friend interface {
	SayHello()
}

func Greet(f Friend) {
	f.SayHello()
}

type Cat struct{}

func (c *Cat) SayHello() {
	fmt.Println("Meow")

}
func main() {
	var guy = new(Person)
	guy.Name = "Mino"
	Greet(guy)

	var cat = new(Cat)
	Greet(cat)
}
