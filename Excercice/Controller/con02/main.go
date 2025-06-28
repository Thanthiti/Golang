package main

import (
	"fmt"
	"sort"
)

type Message struct {
	Id   int
	Text string
}
type Action struct{
	Type string
	Message Message
}

type MessageCTLR struct {
	message []Message
	deleted []Message
	history []Action
	nextID  int
}

func NewMessageCTLR() *MessageCTLR {
	return &MessageCTLR{
		message: []Message{},
		deleted: []Message{},
		history: []Action{},
		nextID:  1,
	}
}

func (c *MessageCTLR) AddMessage(text string) {
	msg := Message{Id: c.nextID, Text: text}
	c.message = append(c.message, msg)
	c.history = append(c.history,Action{Type:"Add",Message: msg })
	c.nextID++
	fmt.Println("Add Success!")
}

func (c *MessageCTLR) RemoveLast() {
	if len(c.message) == 0{
		return
	}
	last := c.message[len(c.message)-1]
	
	c.deleted = append(c.deleted, last)
	c.history = append(c.history, Action{Type:"Remove",Message: last})
	c.message = c.message[:len(c.message)-1]
	fmt.Println("Remove Success!")
}

// UnDo Remove & Undo Add
func (c *MessageCTLR) Undo() {
	if len(c.message) == 0{
		return
	}
	
	lastAction := c.history[len(c.history)-1]
	c.history =  c.history[:len(c.history)-1]

	switch lastAction.Type{
	case "Add":
		if len(c.message) > 0{
			c.message = c.message[:len(c.message)-1]
			fmt.Println(c.message)
		}
	case "Remove":
		c.message = append(c.message, lastAction.Message)
		if len(c.deleted) > 0 {
			c.deleted = c.deleted[:len(c.deleted)-1]
		}
	}
	
	sort.Slice(c.message, func(i, j int) bool {
		return c.message[i].Id < c.message[j].Id
	})

	fmt.Println("Undo Success!")
	
}

func (c *MessageCTLR) ShowMassages() {
	fmt.Println("\n--- Message ---")
	for _,n := range c.message{
		fmt.Println(n)
	}
	fmt.Println("--- Message ---\n")
}
func (c *MessageCTLR) ShowDeleteMassages() {
	fmt.Println("\n--- Deleted ---")
	for _,n := range c.deleted{
		fmt.Println(n)
	}
	fmt.Println("--- Deleted ---\n")
}

func (c *MessageCTLR) ShowHistoryMassages() {
	fmt.Println("\n--- History ---")
	for _,n := range c.history{
		fmt.Println(n)
	}
	fmt.Println("--- History ---\n")
}

func main() {
	ctlr := NewMessageCTLR()
	ctlr.AddMessage("Thanthiti")
	ctlr.AddMessage("Pinmoleep")
	ctlr.RemoveLast()
	ctlr.ShowMassages()
	ctlr.AddMessage("Palm")
	ctlr.ShowMassages()
	ctlr.ShowHistoryMassages()
	ctlr.Undo()
	ctlr.Undo()
	ctlr.ShowMassages()
	ctlr.ShowHistoryMassages()

	// ctlr.ShowHistoryMassages()
	// ctlr.AddMessage("Test")
	// ctlr.Undo()
	// ctlr.ShowMassages()

}