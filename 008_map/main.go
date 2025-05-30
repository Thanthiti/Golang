package main

import "fmt"



func main() {
	name := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}	

	ShowMap(name)
	{
		m := map[string]int{
			"todd":  42,
			"henry": 16,
		}

		if v, ok := m["Padget"]; ok {
			fmt.Printf("%s is %d years old\n", "Padget", v)
		} else {
			fmt.Printf("%s not found\n", "Padget")
		}

		//delete
		m["Shakespeare"] = 459
		fmt.Printf("%#v\n", m)
		delete(m, "Shakespeare")
		fmt.Printf("%#v\n", m)
		delete(m, "Shakespeare") // no pname2ic

		// len
		fmt.Println("len: ", len(m))
	}

	{
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

