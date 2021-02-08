package main

import "fmt"

func main() {
	c := Circle{radius: 10}
	r := Rectangle{length: 10, width: 20}

	fmt.Println(c.area())
	fmt.Println(r.area())

	shapes := []Shape{c, r}
	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (c Circle) area() float64 {
	return 3.141 * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}
