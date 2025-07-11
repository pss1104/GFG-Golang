package main

import "fmt"

func main() {
	//map declaration
	var em map[string]int
	fmt.Println("empty map: ", em)

	//using map literal
	m := map[string]int{
		"banana": 6,
		"apple":  10,
		"mango":  20, //comma needed for last k-v pair also
	}
	fmt.Println("fruit-cost map: ", m)

	//using make to create a map (multi dimensional map)
	mm := make(map[string]map[string]int)
	mm["Sai"] = map[string]int{
		"Maths":   100,
		"Science": 99,
	}
	mm["Niyathi"] = map[string]int{
		"Maths":   90,
		"Science": 100,
	}
	fmt.Println("marks-map: ", mm)

	//deletion from fruit-cost map
	delete(m, "banana")
	fmt.Println("Deleted banana from fruit map, map: ", m)

	//checking if key exists
	val, exists := mm["Sai"]["Maths"]
	if exists {
		fmt.Println("Sai->Maths exists, value is: ", val)
	} else {
		fmt.Println("Sai->Maths does not exist")
	}
}
