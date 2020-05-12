package main

import "fmt"

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
	for i := 1; i <= 1000; i++ {
		fmt.Println(i)
	}
}

func exercise2() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		fmt.Printf("\t%#U\n", i)
		fmt.Printf("\t%#U\n", i)
		fmt.Printf("\t%#U\n", i)
	}
}

func exercise3() {
	birth := 1995
	today := 2020
	for birth <= today {
		fmt.Println(birth)
		birth++
	}
}

func exercise4() {
	birth := 1995
	for {
		fmt.Println(birth)
		if birth == 2020 {
			break
		}
		birth++
	}
}

func exercise5() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}

func exercise6() {
	a := 50
	b := 250
	if a < b {
		fmt.Printf("%d < %d\n", a, b)
	}
}

func exercise7() {
	a := 250
	b := 50

	if a < b {
		fmt.Printf("%d < %d\n", a, b)
	} else if a > b {
		fmt.Printf("%d > %d\n", a, b)
	} else {
		fmt.Printf("%d == %d\n", a, b)
	}
}

func exercise8() {

	a := 250
	b := 250

	switch {
	case a < b:
		fmt.Printf("%d < %d\n", a, b)

	case a > b:
		fmt.Printf("%d > %d\n", a, b)

	case a == b:
		fmt.Printf("%d == %d\n", a, b)

	}
}

func exercise9() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}
}

func exercise10() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
