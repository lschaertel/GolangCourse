package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
}

func exercise1() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("This is the first go routine ")
		wg.Done()
	}()
	go func() {
		fmt.Println("This is the second go routine ")
		wg.Done()
	}()
	wg.Wait()
}

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak()
}

func (p *person) speak() {

}

func exercise2() {
	p1 := person{
		first: "First",
		last:  "Second",
		age:   32,
	}
	saySomething(&p1)
	//saySomething(p1) Doesnt run - expects a pointer
}

func saySomething(h human) {
	h.speak()
}

func exercise3() {
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			rv := counter
			runtime.Gosched()
			rv++
			counter = rv
			wg.Done()
		}()
	}
	fmt.Println("Counter:\t", counter)
	wg.Wait()
}

func exercise4() {
	var wg sync.WaitGroup
	counter := 0
	var mtx sync.Mutex
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mtx.Lock()
			rv := counter
			rv++
			counter = rv
			mtx.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:\t", counter)
}

func exercise5() {
	var wg sync.WaitGroup
	var counter int64 = 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:\t", counter)
}

func exercise6() {
	fmt.Println("GOOS:\t", runtime.GOOS)
	fmt.Println("GOARCH:\t", runtime.GOARCH)
}
