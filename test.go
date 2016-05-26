package main

import (
	//"time"
	"fmt"
	//"strconv"
	//"github.com/angeldm/euler/problem"
	"github.com/angeldm/euler/problem"
)

func main() {
	fmt.Println(problem.P004())

	//fmt.Println(goodProd(12, 0, 12))

}
func goodProd(n int, min int, max int) (bool, []int) {
	//fmt.Println("Checking prod")
	var fac []int
	for i := 5; i <= max; i++ {
		if n % i == 0 {
			fmt.Println(i)
			if n / i > min {
				fmt.Println(n/i)
				fac = append(fac,i)
				fac = append(fac, n/i)
				return true, fac
			}
		}
	}
	return false, fac
}