package main

import "fmt"

func main() {

	exercise1()
	exercise23()
	exercise4()
	exercise5()
	exercise6()
	exercise7()
	exercise10()
}

func exercise1() {
	s := [5]int{1, 2, 3, 4, 5}

	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", s)
}

func exercise23() {
	s := [10]int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", s)

	s1 := append(s[:5])
	fmt.Println(s1)
	s2 := append(s[5:])
	fmt.Println(s2)
	s3 := append(s[2:7])
	fmt.Println(s3)
	s4 := append(s[1:6])
	fmt.Println(s4)
}

func exercise4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func exercise5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)
	fmt.Println(y)

}

func exercise6() {
	x := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, `California`, ` Colorado`, ` Connecticut`,
		` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`,
		` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`,
		` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`,
		` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`,
		` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`,
		` West Virginia`, ` Wisconsin`, ` Wyoming`,}

	length := len(x)
	capacity := cap(x)
	fmt.Printf("Length '%v' Capacity '%v'", length, capacity)

	for i := 0; i < length; i++ {
		fmt.Println(i, x[i])
	}
}

func exercise7() {
	s := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}

	for i, v := range s {
		fmt.Println(i, v)
		for i2, v2 := range v {
			fmt.Println(i2, v2)
		}
	}
}

func exercise10() {
	m := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},}

	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	delete(m, "moneypenny_miss")
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
