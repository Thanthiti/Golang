package main
import "fmt"
func main(){

	a := [3]int{42, 43, 44}
	fmt.Println(a)

	b := [...]string{"Hello", "World!","Welcome To golang"}
	fmt.Println(b)

	var c [2]int
	c[0] = 7
	c[1] = 8
	fmt.Println(c)

	{
		var ni [7]int
		ni[0] = 42
		fmt.Printf("%#v \t\t\t\t %T\n", ni, ni)

		ni2 := [4]int{55, 56, 57, 58}
		fmt.Printf("%#v \t\t\t\t\t %T\n", ni2, ni2)

		ns := [...]string{"Chocolate", "Vanilla", "Strawberry"}
		fmt.Printf("%#v \t %T\n", ns, ns)

		fmt.Println(len(ni))
		fmt.Println(len(ni2))
		fmt.Println(len(ns))
	}


	Sli := []int{1, 2}
	fmt.Println(Sli)
	Sli = append(Sli, 3,4,5,6)
	fmt.Println(Sli, " " , len(Sli))

	// Multi dimension
	{

	a2 := [][]int{{1,2},{3,4}}
	fmt.Println("Array 2 Dimension : ",a2)

	a2[0] = append(a2[0],1,3,4 )
	fmt.Println("Array 2 Dimension : ",a2)

	// Add New Row 
	a2 = append(a2,[]int{5,4} )
	fmt.Println("Array 2 Dimension : ",a2)

	// Add New Row with loop
	for i := 0; i < 3; i++ {
		newRow := []int{i, i + 10} 
		a2 = append(a2, newRow)
	}
	fmt.Println("Array 2 Dimension : ",a2)

	
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Guiness", "Wolverine Tracks"}
	fmt.Println(jb)
	fmt.Println(jm)

	xxs := [][]string{jb, jm}
	fmt.Println(xxs)
	}
}

