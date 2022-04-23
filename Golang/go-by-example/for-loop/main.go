// For Loop

package main

import "fmt"

func main() {
	i := 1 // assigning the value of variable (type: int)

	// For loop is used to iterate the statement
	for i <= 5 { // i is greater or less than 5
		fmt.Println(i) // print the the value of the var i
		i = i + 1      // add + 1, So if the value of var i is 3, then it will add + 1.
	}

	for j := 8; j <= 12; j++ { // ++ is used to increment by 1
		fmt.Println(j)
	}

	for {
		fmt.Println("loop") // if we don't give break, then it will iterate forever.
		break               // break immediately terminate the loop.
	}

	for n := 1; n <= 8; n++ {
		if n%2 == 0 { // % is used for modulus
			continue // continue is used for skipping the true statement
		}
		fmt.Println(n)
	}
}
