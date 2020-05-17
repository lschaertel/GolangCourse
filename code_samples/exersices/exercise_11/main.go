package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
}

type person struct {
	First   string
	Last    string
	Sayings []string
}

func exercise1() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}

func exercise2() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("coudn't marshal this object %v", a)
	}
	return bs, nil
}

func exercise3() {
	err := customErr{"This is a custom error message"}
	foo(err)

}

type customErr struct {
	err string
}

func (c customErr) Error() string {
	return fmt.Sprintln(c.err)
}

func foo(err error) {
	fmt.Print(err)
}

func exercise4() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		err := sqrtError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  fmt.Errorf("the passed integer cant be below 0"),
		}
		return 0, err
	}
	return 42, nil
}

func exercise5() {

}
