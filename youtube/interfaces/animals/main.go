package main

import "fmt"

type animal interface {
	makeSound()
}

type cat struct{}
type dog struct{}

func (c *cat) makeSound() {
	fmt.Println("meow!")
}

func (d *dog) makeSound() {
	fmt.Println("woof!")
}

func main() {
	var c, d = &cat{}, &dog{}
	c.makeSound()
	d.makeSound()
}
