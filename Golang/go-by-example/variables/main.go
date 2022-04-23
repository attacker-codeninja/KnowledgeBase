// Variables

package main

import "fmt"

func main() {
	var a = "abcd" // we can use single variable and give its value
	fmt.Println(a)

	var b, c int = 2, 3 // we can also use multiple variables with data types
	fmt.Println(b, c)
	fmt.Println(b * c) // multiplication
	fmt.Println(b + c) // addition

	var d = true // we can also use booleans
	fmt.Println(d)

	var e int
	fmt.Println(e) // this will print 0, because we haven't gave value for this

	f := "sarin" // this is the shorthand to write the variable.
	fmt.Println(f)
}
