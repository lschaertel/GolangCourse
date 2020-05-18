package main

import (
	"GolangCourse/code_samples/exersices/exercise_13/ex1/dog"
	"GolangCourse/code_samples/exersices/exercise_13/ex2/quote"
	"GolangCourse/code_samples/exersices/exercise_13/ex2/word"
	"GolangCourse/code_samples/exersices/exercise_13/ex3/mymath"
	"fmt"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
}

type canine struct {
	name string
	age  int
}

func exercise1() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}

func exercise2() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}

}

func exercise3() {
	xxi := gen()
	for _, v := range xxi {
		fmt.Println(mymath.CenteredAvg(v))
	}

}

// gen generates data to pass into CenteredAvg
func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}
