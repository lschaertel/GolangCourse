package main

import (
	"fmt"
	"math"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
	exercise7()
	exercise8()
	exercise9()
	exercise10()
}

func exercise1() {
	i := foo1()
	i2, s := bar1()
	fmt.Println(i)
	fmt.Println(i2, s)
}

func foo1() int {
	return 42
}

func bar1() (int, string) {
	return 42, "420"
}

func exercise2() {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := foo2(i...)
	sum2 := bar2(i)
	fmt.Println(sum, sum2)
}

func foo2(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func bar2(i []int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func exercise3() {
	fd := func() {
		fmt.Println("This is a defered function call")
	}
	defer fd()
	fmt.Println("This is a normal function call")
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("This is a person speaking", p)
}

func exercise4() {
	p := person{
		first: "Ludwig",
		last:  "Schärtel",
		age:   24,
	}
	p.speak()
}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (s square) area() float64 {
	return s.length * s.length
}

func info(s shape) {
	fmt.Println(s.area())
}

func exercise5() {
	circ := circle{
		radius: 12.345,
	}

	squa := square{
		length: 15,
	}

	info(circ)
	info(squa)
}

func exercise6() {
	func() {
		fmt.Println("This is an anonymous function call")
	}()

}

func exercise7() {
	f := func() {
		fmt.Println("This variable contains an anonymous function to call")
	}
	f()
}

func exercise8() {
	f := foo()
	f()
}

func foo() func() {
	return func() {
		fmt.Println("This function got returned from an function")
	}
}

func exercise9() {
	funcPass(foo())
}

func funcPass(f func()) {
	f()
}

func exercise10() {
	{
		i := 5
		fmt.Println(i)
	}
	{
		i := 10
		fmt.Println(i)
	}
}
