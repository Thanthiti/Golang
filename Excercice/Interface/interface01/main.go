// Cal Triangle by interface

package main

import (
	"fmt"
)

type shaper interface{
	calArea() float64
	calPerimeter() float64
}
type triangle struct{
	width ,height ,base float64
}
func (t triangle) calArea() float64{
	return 0.5 *t.base*t.height
}
func (t triangle) calPerimeter() float64{
	return t.base + t.height +t.width
}

func showPerimeter (s shaper) float64{
	return s.calPerimeter()
}
func showArea (s shaper) float64{
	return s.calArea()
}


func main(){
	triangle := triangle{
		width: 3,
		height: 4,
		base: 6,
	}
	perimeter := showPerimeter(triangle)
	area := showArea(triangle)
	
	fmt.Printf("Show Perimeter : %.2f\n",perimeter)
	fmt.Printf("Show Area : %.2f",area)

}