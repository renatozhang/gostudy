package main

import "fmt"

func main() {
	/*
		const a int = 100
		const b string = "hello"
		// a = 200 //报错 常量不能修改
	*/
	/*
		const (
			a int = 100
			b
			c int = 200
			d
		)
	*/
	const (
		a = 100
		b
		c = 200
		d
	)

	fmt.Printf("a-%d b-%d c-%d d-%d\n", a, b, c, d)
	/*
		const (
			e = iota
			f = iota
			g = iota
		)
	*/
	/*
		const (
			e = iota
			f
			g
		)
	*/
	// fmt.Printf("e-%d f-%d g-%d\n", e, f, g)
	const (
		a1 = 1 << iota
		a2
		a3
		a4
	)

	fmt.Printf("a1-%d a2-%d a3-%d a4-%d\n", a1, a2, a3, a4)

}
