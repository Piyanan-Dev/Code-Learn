package main

import "fmt"

type Rect struct{ W, H float64 }

func (r Rect) Area() float64      { return r.W * r.H }
func (r Rect) Paramiter() float64 { return (r.W + r.H) * 2 }

func main() {

	sqaur := []Rect{
		{W: 2, H: 3},
		{W: 4, H: 5},
	}
	for _, v := range sqaur {
		fmt.Println(v.Area(), " ", v.Paramiter())
	}

}
