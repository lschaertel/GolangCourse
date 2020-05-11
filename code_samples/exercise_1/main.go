package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise45()
	exercise6()
}

func exercise1() {
	a := 42
	b := "James Bond"
	c := true

	fmt.Println(a, b, c)
	fmt.Println("Value of a: ", a)
	fmt.Println("Value of b: ", b)
	fmt.Println("Value of c: ", c)
}

var x int
var y string
var z bool

func exercise2() {
	fmt.Println("Value of x: ", x)
	fmt.Println("Value of y: ", y)
	fmt.Println("Value of z: ", z)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}

func exercise3() {
	x = 42
	y = "James Bond"
	z = true

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}


type intDefault int
func exercise45() {
	var x intDefault

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}

func exercise6() {

}
