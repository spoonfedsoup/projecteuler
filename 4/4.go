package main

import (
	"strconv" //package to convert between numeric data types and strings.
	"fmt"
)

func main() {
	min := 100*100
	max := 999*999
	var fac1, fac2 int //declares fac1 and fac2 as ints

	for i := max; i >= min; i--{
		s := strconv.Itoa(i) // converts int to string

		if isPalindrome(s){
			n := i
			flag, n, m := goodProd(n, 100, 999) //the blank identifier _ signals a return will not be used
			if flag {
				fac1 = n//so the only value being passed is a boolean value
				fac2 = m
				break
			}
		}

	}
	fmt.Println(fac1*fac2," = ",fac1," * ", fac2)
}
func isPalindrome(s string) bool{ // checks for palindrome
	last := len(s)-1

	for i:=0; i<=last/2; i++{
		if s[i]!=s[last-i]{ //a string is actually a slice in Go
			return false
		}
	}
	return true
}
func goodProd(n int, min int, max int) (bool, int , int) { // looks for a pair of divisors of n between min and max
	for i := min; i <= max; i++ {                      // the (bool,int,int) declares function will return 3 objects
		if n % i == 0 {
			m := n/i
			if m > min && m <= max {
				return true, i, m
			}
		}
	}
	return false, 1, 1
}
