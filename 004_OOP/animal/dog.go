package animal
import "fmt"

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println(d.Name, "says: Woof!")
}