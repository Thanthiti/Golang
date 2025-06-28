package main

import (
	"fmt"
)


type Number struct {
	Value int
}

type NumberController struct {
	counter Number
	History []int
}

func NewNumberController() *NumberController {
	return &NumberController{
		counter: Number{Value: 0},
	}
}

func (c *NumberController) IncreaseNumber() {
	c.counter.Value++
	c.History = append(c.History,c.counter.Value)
}

func (c *NumberController) DecreaseNumber() {
	c.counter.Value--
	c.History = append(c.History,c.counter.Value)
}

func (c *NumberController) ResetNumber() {
	c.counter.Value = 0
}

func (c *NumberController) SetNuber(number int) {
	c.counter.Value = number
	c.History = append(c.History,c.counter.Value)
}

func (c *NumberController) Undo() {
	if len(c.History) == 0{
		return
	}
		
		c.History = c.History[:len(c.History)-1]
		
		if len(c.History) > 0{
			c.counter.Value = c.History[len(c.History)-1]
			}else{
				c.counter.Value = 0
			}
	}

func (c *NumberController) ShowNumber() {
	fmt.Println("Current number is : ",c.counter.Value )
}
func (c *NumberController) ShowHistory() {
	fmt.Println("History number is : "	,c.History )
}

func main() {
	ctl := NewNumberController()
	ctl.IncreaseNumber()	
	ctl.IncreaseNumber()	

	ctl.ShowNumber()	
	ctl.ShowHistory()	
	
	ctl.SetNuber(10)
	ctl.ShowNumber()	
	ctl.SetNuber(15)
	
	ctl.IncreaseNumber()
	ctl.ShowHistory()	
	ctl.ShowNumber()	
	
	ctl.Undo()
	ctl.ShowNumber()	
	ctl.ShowHistory()	
	
}