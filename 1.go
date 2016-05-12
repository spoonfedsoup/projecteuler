package main // Root file designation

import (
	"fmt" //fmt is an output package
)

func main() { //Where root file executes code

	var mult []int //Slice declaration. A slice is a dynamic array

	for i := 1; i < 1000; i++ { //all loops are for loops in Go
		if i%3 == 0 || i%5 == 0 {
			mult = append(mult, i)
		}
	}

	sum := 0

	for i := 0; i<len(mult); i++ {
		sum += mult[i]
	}

	fmt.Println(sum)
}

/* There are two ways to declare a variable in Go.
Shorthand is done by := and is typed automatically.
So a := 0 is declaring an integer variable a and setting it to 0.
Second is var x type.
Example:
var a int
a = 0
first it set aside space for an int-sized a, then assigned it to 0.

Lastly, the people who made Go found a way to reduce all loops to a
for loop.
 */
