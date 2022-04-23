// Constants

package main

import (
	"fmt"
	"math" // imported math package
)

const s string = "shreya" // constant is a constant value

func main() {
	fmt.Println(s)

	const n = 500000 // this will assign an constant integer
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))    // we can also define types on constant
	fmt.Println(math.Sin(n)) // math operations

}
