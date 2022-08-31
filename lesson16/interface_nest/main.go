package main

import "fmt"

type Animal interface {
	Eat()
	Talk()
	Name() string
}

type Describe interface {
	Describle() string
}

type AdvanceAnimal interface {
	Animal
	Describe
}

type Dog struct{}

func (d Dog) Eat() {
	fmt.Println("dog is eating")
}
func (d Dog) Talk() {
	fmt.Println("dog is talking")
}
func (d Dog) Name() string {
	fmt.Println("my name is dog")
	return "dog"
}

func (d Dog) Describle() string {
	fmt.Println("dog is a dog")
	return "dog is dog"
}

func main() {
	var d Dog
	var a AdvanceAnimal
	a = d
	a.Describle()
	a.Eat()
	a.Talk()
	a.Name()
}
