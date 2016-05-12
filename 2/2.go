package main

import "fmt"

func main() { 

	p := 0
	q := 1
	sum := 0

	for i := 0; p <= 4000000; i++ { //calculates nth Fib
		p, q = q, p+q
		if p%2==0{   //checks for eveness
			sum += p  //adds if even
		}
	}
	fmt.Println(sum)
}
//I wanted to use the closed-form Binet's Formula to calculate
//the nth fibonnaci number, but type casting was cumbersome.
//The complexity of recursion is much higher, but it's kind of cool.