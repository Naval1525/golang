package main

import "time"

func main() {
	// type 1
	i := 5
	switch i {
	case 1:
		println("i is 1")
	case 2:
		println("i is 2")
	case 3:
		println("i is 3")
	case 4:
		println("i is 4")
	case 5:
		println("i is 5")
	default:
		println("i is not 1, 2, 3, 4, or 5")
	}

	//multiple switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("It's the weekend!")
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		println("It's a weekday!")
	default:
		println("Unknown day")
	}

	//type switch
	whoAmI := func(i interface{}) {
		switch i := i.(type) {
		case int:
			println("I am an int:", i)
		case string:
			println("I am a string:", i)
		case bool:
			println("I am a bool:", i)
		default:
			println("I am of a different type:", i)
		}
	}
	whoAmI(42)
}
