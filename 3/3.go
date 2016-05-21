package main

import (
	"fmt"
	"github.com/angeldm/euler/util/primegen"
)	//use go get in command prompt to download package
	//files into workspace, otherwise import can't find
	//package and throws and error.


func main() {
	var n uint64 = 600851475143 //primegen outputs uint64
	var prod uint64 = 1 //holds product of prime diviors of n
	var largestPrime uint64
	pg := primegen.New() //pg is the prime generator object

	for i := pg.Next() ; i <= n/2; i = pg.Next() {
		if divides(i, n) {
			largestPrime = i
			prod *= i
		}
		if prod > n/2{ // n/2 is largest prime that
			break  // could divide n since 2 is
		}              // the smallest prime that
	}                      // could divide n.
	fmt.Println(largestPrime)
}

func divides(n uint64, m uint64) bool { //n divides m
	flag := false

	if m%n == 0 {
		flag = true
	}
	return flag
}
