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
}

func exercise1() {
	c := make(chan int)

	go func() { c <- 42 }()

	fmt.Println(<-c)

	c2 := make(chan int, 1)

	c2 <- 42

	fmt.Println(<-c2)
}
func exercise2() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)

}

func exercise3() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
func exercise4() {
	q := make(chan int)
	c := genSelect(q)

	receiveSelect(c, q)

	fmt.Println("about to exit")
}

func receiveSelect(c1, c2 <-chan int) {
	for {
		select {
		case v := <-c1:
			fmt.Println(v)
		case <-c2:
			return
		}
	}
}

func genSelect(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func exercise5() {
	c := make(chan int)
	go func() { c <- 42 }()
	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)

}

func exercise6() {
	c1 := make(chan int)
	send(c1)

	for v := range c1 {
		fmt.Println(v)
	}
}

func send(c chan<- int) {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
}

func exercise7() {
	fmt.Println("Exercise 7")
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go send7(c)
	}

	for v := range c {
		fmt.Println(v)
	}
}

func send7(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
