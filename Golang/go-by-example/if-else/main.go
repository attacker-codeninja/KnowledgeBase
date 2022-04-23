// If or Else Statement

package main

import "fmt"

func main() {
	if 8%2 == 0 { // if 8 divides 2 and the reminder is 0 then print below statement
		fmt.Println("8 is even!")
	} else {
		fmt.Println("8 is odd!") // or if the reminder is 1 or more then print below statement
	}

	if 9%3 == 0 {
		fmt.Println("9 is divisible by 3.")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative!")
	} else if num < 10 { // we can also use nested else if statement
		fmt.Println(num, "has 1 digit.")
	} else {
		fmt.Println(num, "has multiple digits.")
	}
}
