package main

import "fmt"
import "log"

func gcd(a, b int) int {
	//log.Println("gcd: ", a, b)
	for a % b != 0 {
		tmp := a % b
		a = b
		b = tmp
		//log.Println(a, b)
	}
	return b
}

func main() {
	var terms []int
	terms = append(terms, 2)

	for i := int(1); i < 40; i++ {
		terms = append(terms, 1, 2*i, 1)
	}

	N := 9

	numer := terms[N]
	denom := 1

	for i := N - 1; i >= 0; i-- {
		log.Println(numer, " / ", denom)

		new_numer := terms[i] * numer + denom
		new_denom := numer

		_gcd := gcd(numer, denom)

		log.Println("new_numer: ", new_numer, " new_denom: ", new_denom, " _gcd: ", _gcd)

		numer = new_numer / _gcd
		denom = new_denom / _gcd
	}
	fmt.Println(numer, " / ", denom)
}
