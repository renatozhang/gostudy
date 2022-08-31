package main

import "fmt"

type Employer interface {
	CalcSalary() float32
}

type Programer struct {
	name  string
	base  float32
	extra float32
}

func NewProgramer(name string, base float32, extra float32) Programer {
	return Programer{
		name:  name,
		base:  base,
		extra: extra,
	}
}

func (p Programer) CalcSalary() float32 {
	return p.base
}

type Sale struct {
	name  string
	base  float32
	extra float32
}

func NewSale(name string, base float32, extra float32) Sale {
	return Sale{
		name:  name,
		base:  base,
		extra: extra,
	}
}

func (p Sale) CalcSalary() float32 {
	return p.base + p.extra*p.base*0.5
}

func calcAll(e []Employer) float32 {
	var cost float32
	for _, employer := range e {
		cost += employer.CalcSalary()
	}
	return cost
}
func main() {
	p1 := NewProgramer("搬砖1", 1500, 0)
	p2 := NewProgramer("搬砖2", 1500, 0)
	p3 := NewProgramer("搬砖3", 1500, 0)

	s1 := NewSale("销售1", 800, 2.5)
	s2 := NewSale("销售2", 800, 2.5)
	s3 := NewSale("销售3", 800, 2.5)

	var employList []Employer
	employList = append(employList, p1)
	employList = append(employList, p2)
	employList = append(employList, p3)

	employList = append(employList, s1)
	employList = append(employList, s2)
	employList = append(employList, s3)

	cost := calcAll(employList)
	fmt.Printf("这个月的总人力成本：%f\n", cost)

}
