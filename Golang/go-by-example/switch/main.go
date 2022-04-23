// Switch Statement

package main

import (
	"fmt"
	"time" // imported time package
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ") // Print
	switch i {                     // switch is used to switch the statement a/q the case
	case 1: // here we used case 1
		fmt.Println("one")
	case 2: // value of var i is 2 and if the case matches with this,
		fmt.Println("two") // then print this.
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() { // we can also use switch directly on func
	case time.Saturday, time.Sunday: // if the day is Saturday/Sunday then print below statement.
		fmt.Println("Weekend Day, Go and Get drink!")
	default: // Or if the day is other than Sat/Sun then print below statement.
		fmt.Println("Oh, Snap! It's Weekday.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Breakfast time!")
	default:
		fmt.Println("Oh! It's after noon.")
	}

	whatAmI := func(i interface{}) { // here we used interface{}
		switch t := i.(type) { // type statement is used to denote the data type of given interfaces.
		case bool:
			fmt.Println("I'm a boolean.")
		case int:
			fmt.Println("I'm a integer!")
		default:
			fmt.Printf("Oaaaah type is: %T\n", t)
		}
	}

	whatAmI(true) // we are calling the whatAmI function which we created in upper section
	whatAmI(3)
	whatAmI("Hello!")
}
