package main

import "fmt"
import "math"

type geometry interface {
	area() float64
	perim() float64
}

type rect1 struct {
	width, height float64
}

type circle struct {
	raduis float64
}

func (r rect1) area() float64  {
	return r.width * r.height
}

func (r rect1) perim() float64  {
	return 2 * r.width + 2* r.height
}

func (c circle) area() float64  {
	return math.Pi * c.raduis * c.raduis
}


func (c circle) perim() float64  {
	return 2 * math.Pi * c.raduis
}

func measure(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.area())
}


func main() {

	r := rect1{width:3, height:4}
	c := circle{raduis:5}

	measure(r)
	measure(c)
}