package main

import "fmt"

func main() {
	age := 18
	if age >= 18 {
		fmt.Println("adult")
	} else {
		fmt.Println("not adult")
	}

	var role = "admin"
	var hasPermissions = true
	if role == "admin" && hasPermissions {
		fmt.Println("Access granted")
	} else {
		fmt.Println("Access denied")
	}

	if name := "naval"; name == "naval" {
		fmt.Println("Hello Naval")
	} else {
		fmt.Println("Hello Unknown")
	}

	// go doesnt have ternary, for current version

}
