package main

func main() {
	// range is used to iterate over arrays, slices, maps, strings, and channels
	// it returns two values: the index (or key) and the value at that index (or key)

	// Example with an array
	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		println("Index:", i, "Value:", v)
	}

	// Example with a slice
	slice := []string{"apple", "banana", "cherry"}
	for i, v := range slice {
		println("Index:", i, "Value:", v)
	}

	// Example with a map
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		println("Key:", k, "Value:", v)
	}

	// Example with a string
	str := "hello"
	for i, v := range str {
		println("Index:", i, "Rune Value:", string(v))
	}
}
