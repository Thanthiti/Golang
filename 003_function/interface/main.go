// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Abser interface {
// 	Abs() float64
// }

// func main() {
// 	var a Abser
// 	f := MyFloat(-math.Sqrt2)
// 	v := Vertex{3, 4}

// 	a = f  // a MyFloat implements Abser
// 	a = &v // a *Vertex implements Abser

// 	// In the following line, v is a Vertex (not *Vertex)
// 	// and does NOT implement Abser.
// 	a = &v

// 	fmt.Println(a.Abs())
// }

// type MyFloat float64

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// type Vertex struct {
// 	X, Y float64
// }

// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }
package main

import "fmt"

type Shaper interface {
	Area() float64
	Perimeter() string 
}

type Rectangle struct {
	width  float64
	height float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.width * rectangle.height
}
func (r Rectangle) Perimeter() string {
    return "Rectangle Perimeter\n"
}

type Square struct {
	width float64
}

func (square Square) Area() float64 {
	return square.width * square.width
}
func (s Square) Perimeter() string {
	return "Square Perimeter\n"
}

type Circle struct {
	redius float64
}

func (circle Circle) Area() float64 {
	return circle.redius * circle.redius * 3.1415
}
func (c Circle) Perimeter() string {
	return "Circle Perimeter\n"
}

func ComputeArea(shape Shaper) {
	// fmt.Printf("Area of %+v=%.3f\n", shape, shape.Area())
}
func ComputeParameter(shape Shaper) {
	fmt.Printf("Perimeter of %+v=%s\n", shape, shape.Perimeter())
	fmt.Printf("Area of %+v=%.3f\n", shape, shape.Area())
}

func main() {
	rectangle := Rectangle{5, 10}
	square := Square{5}
	circle := Circle{5}
	// fmt.Println(rectangle)
	
	shapes := [...]Shaper{rectangle, square, circle}
	for _, shape := range shapes {
		ComputeParameter(shape)
		ComputeArea(shape)
	}


}