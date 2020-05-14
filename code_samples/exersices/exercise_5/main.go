package main

import "fmt"

type person struct {
	first       string
	last        string
	favoriteIce []string
}

func main() {
	exercise12()
	exercise3()
	exercise4()
}

func exercise12() {
	p1 := person{
		first:       "Ludwig",
		last:        "Schärtel",
		favoriteIce: []string{"Chocolate", "Vanilla"},
	}
	p2 := person{
		first:       "Sudwig",
		last:        "Lärtel",
		favoriteIce: []string{"Mint", "Strawberry"},
	}
	fmt.Println(p1.first, p1.last)
	for _, v := range p1.favoriteIce {
		fmt.Printf("\t%v\n", v)
	}
	fmt.Println(p2.first, p2.last)
	for _, v := range p2.favoriteIce {
		fmt.Printf("\t%v\n", v)
	}
	fmt.Println()

	m := map[string]person{p1.last: p1, p2.last: p2}
	fmt.Println(m)
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func exercise3() {
	t1, t2 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "Black",
		},
		fourWheel: true,
	}, sedan{
		vehicle: vehicle{
			doors: 4,
			color: "White",
		},
		luxury: true,
	}

	fmt.Println(t1)
	fmt.Println(t1.doors, t1.color, t1.fourWheel)
	fmt.Println(t2)
	fmt.Println(t2.doors, t2.color, t2.luxury)
}
func exercise4() {

	x := struct {
		name string
		location string
		age int
	}{
		name: "Gopher",
		location: "Earth",
		age: 10,
	}
	fmt.Println(x)
}
