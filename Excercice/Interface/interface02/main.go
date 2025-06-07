// unit Converter bt interface
package main

import "fmt"

type Converter interface{
	Converter() float64
}

type MeterToCentimeter struct{
	value float64
}

func (m MeterToCentimeter) Converter() float64{
	return m.value * 100
}

type KilogramToGram struct{
	value float64
}

func (k KilogramToGram) Converter() float64{
	return k.value * 1000
}

func ShowConversion(c Converter){
	fmt.Printf("Converted value : %.2f",c.Converter() )
	
}
func main(){
	meter := MeterToCentimeter{5}
	kilo := KilogramToGram{2}
	
	ShowConversion(meter)
	ShowConversion(kilo)
}




