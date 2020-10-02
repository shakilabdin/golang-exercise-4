package main

import (
	"fmt"
)

func main() {
	// exercise_one()
	exercise_two()
	exercise_three()
	exercise_four()
	exercise_five()
}

func exercise_one()  {
	for i := 0; i <= 10000; i++ {
		fmt.Println(i)
	}
}
func exercise_two()  {
	bd := 1985
	for bd <= 2017 {
		fmt.Println(bd)
		bd++
	}
}
func exercise_three()  {
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the remainder (aka modulus) is %v\n", i, i%4)
	}
}
func exercise_four()  {
	x := "James Bond"
	
	if x == "James Bond" {
		fmt.Println(x)
	}
}
func exercise_five()  {
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}
}
