package main

import "fmt"

type Integer int

func (a Integer) Less (b Integer) bool {
	return a < b
}

type Rect struct {
	x, y float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}


func main () {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a,"less 2")
	}

	rect1 := &Rect{0, 0, 100, 200 }
	fmt.Println("rect1 area is",rect1.Area())
}
