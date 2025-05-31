package main

import (
	"fmt"
	"maps"
)

// maps->hash,object
func main() {
	//creating map
	m := make(map[string]string)
	m["name"] = "golang"
	m["area"] = "programming language"
	//get an element
	fmt.Println(m["name"], m["area"])
	//Imp: if key not found it return the zero value
	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 3}
	fmt.Println(maps.Equal(m1, m2))
}
