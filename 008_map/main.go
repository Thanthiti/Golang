package main

import "fmt"



func main() {
	name := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}	

	ShowMap(name)
	DeleteAndCheckMap(name)
	LoopMap(name)
	

}
func ShowMap(name map[string]int ){

	fmt.Println("The age of Henry was", name["Henry"])

	fmt.Println(name)
	fmt.Printf("%#v\n", name)
	
	name2 := make(map[string]int)
	name2["Lucas"] = 28
	name2["Steph"] = 37
	fmt.Println(name2)
	fmt.Printf("%#v\n", name2)

	fmt.Println(len(name2))
}

func DeleteAndCheckMap(name map[string]int ){

	if v, ok := name["Padget"]; ok {
		fmt.Printf("%s is %d years old\n", "Padget", v)
	} else {
		fmt.Printf("%s not found\n", "Padget")
	}

	//delete
	name["Shakespeare"] = 459
	fmt.Printf("%#v\n", name)
	delete(name, "Shakespeare")
	fmt.Printf("%#v\n", name)
	delete(name, "Shakespeare") 

	// len
	fmt.Println("len: ", len(name))
}

func LoopMap(name map[string]int ){
	fmt.Println("1 : ")
		for k, v := range name {
			fmt.Println(k, v)
		}
		
		fmt.Println("2 : ")
		for _, v := range name {
			fmt.Println(v)
		}
		
		fmt.Println("3 : ")
		for k := range name {
			fmt.Println(k)
		}
}