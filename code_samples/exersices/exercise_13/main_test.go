package main

import (
	"GolangCourse/code_samples/exersices/exercise_13/ex1"
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	n := ex1.Years(10)
	if n != 70 {
		t.Error("got", n, "want", 70)
	}
}

func TestYearsTwo(t *testing.T) {
	n := ex1.YearsTwo(10)
	if n != 70 {
		t.Error("got", n, "want", 70)
	}
}

func ExampleYears() {
	fmt.Println(ex1.Years(10))
	// Output:
	// 70
}

func ExampleYearsTwo() {
	fmt.Println(ex1.YearsTwo(10))
	// Output:
	// 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex1.Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex1.YearsTwo(10)
	}
}
