package main

import "fmt"

func main() {
	exercise1()
	exercise2()
}

func exercise1() {
	i := 42
	fmt.Println(&i)
}

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	(*p).first = "dereferenced"
	(*p).last= "structure"
}

func exercise2() {
	p1 := person{
		first: "Ludwiog",
		last:  "Schärtel",
		age:   24,
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
