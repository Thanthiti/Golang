package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float64
	Perimeter() float64
}
type Shape3D interface {
	Shaper
	Volume() float64
}

type Rectangle struct {
	width, heigth float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.heigth
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.heigth)
}

type Cube struct {
	length float64
}

func (c Cube) Area() float64 {
	return 6 * (c.length * c.length)
}
func (c Cube) Perimeter() float64 {
	return 12 * c.length
}
func (c Cube) Volume() float64 {
	return math.Pow(c.length,3)
}

func PrintShape(s Shaper){
	fmt.Printf("Area: %.2f\n", s.Area())
    fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}
func PrintCube(s3d Shape3D){
	fmt.Printf("Area: %.2f\n", s3d.Area())
    fmt.Printf("Perimeter: %.2f\n", s3d.Perimeter())
    fmt.Printf("Volume: %.2f\n", s3d.Volume())
}

func main() {
	r := Rectangle{5,3}
	c := Cube{4}
	fmt.Println("Rectangle")
	PrintShape(r)
	fmt.Println("Cube")
	PrintCube(c)
}