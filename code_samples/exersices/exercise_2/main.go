package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
}

func exercise1() {
	x := 25

	fmt.Printf("%d\t\t%b\t\t%#X\n", x, x, x)
}
func exercise2() {

	g := 1 == 2
	h := 1 <= 2
	i := 1 >= 2
	j := 1 != 2
	k := 1 < 2
	l := 1 < 2

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

}

const (
	a     = 42
	b int = 43
)

func exercise3() {
	println(a, b)
}

func exercise4() {
	x := 5
	fmt.Printf("%d\t\t%b\t\t%#X\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t\t%b\t%#X\n", y, y, y)
}

func exercise5() {
	var x string = `This is a raw string literal
	over multiple lines`
	fmt.Println(x)
}

const (
	_ = iota

	c = 2020 + iota
	d = 2020 + iota
	e = 2020 + iota
	f = 2020 + iota
)

func exercise6() {

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
