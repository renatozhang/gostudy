package main

import "fmt"

type Animal interface {
	Talk()
	Eat()
	Name() string
}

type PuruDongWu interface {
	TaiSheng()
}
type Dog struct{}

func (d Dog) Talk() {
	fmt.Println("汪汪汪")
}

func (d Dog) Eat() {
	fmt.Println("我在吃骨头")
}

func (d Dog) Name() string {
	fmt.Println("我的名字叫旺财")
	return "旺财"
}

func (d Dog) TaiSheng() {
	fmt.Println("狗是胎生的")
}

func testInterface1() {
	var d Dog
	var a Animal

	fmt.Printf("%v %T %p\n", a, a, a)
	if a == nil {
		// a.Eat()
		fmt.Println("a is nil")
	}

	a = d
	a.Eat()

	var b PuruDongWu
	b = d
	b.TaiSheng()
}

func main() {
	testInterface1()
}
