package main

import "fmt"

type Circle struct {
	x float64
	y float64
	r float64
}

func main() {

	c := Circle{0, 0, 6}
	d := Circle{0, 0, 7}

	if c.r < d.r {
		fmt.Println("d is bigger then c")
	}

}
