package main

func add(a int, b int) int {
	return a + b
}

// multiple return functions
func getlang() (string, string, string) {
	return "Go", "Python", "Java"
}
func main() {
	result := add(5, 3)
	println("The sum is:", result)
	lang1, lang2, lang3 := getlang()
	println("Languages:", lang1, lang2, lang3)

}
