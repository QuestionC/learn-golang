package main

import "fmt"
import "log"
import "math/big"

var big0 big.Int
var big1 big.Int

func gcd(_a, _b *big.Int) *big.Int {
	//log.Println("gcd: ", a, b)
	a := new(big.Int)
	b := new(big.Int)
	c := new(big.Int)

	a.Set(_a)
	b.Set(_b)
	c.Mod(a, b)
	for c.Cmp(big.NewInt(0)) != 0 {
		// log.Println(a, b, c)
		a.Set(b)
		b.Set(c)
		c.Mod(a, b)
		//log.Println(a, b)
	}
	return b
}

func main() {
	var terms []*big.Int
	terms = append(terms, big.NewInt(2))

	for i := int64(1); i < 40; i++ {
		terms = append(terms, big.NewInt(1), big.NewInt(2*i), big.NewInt(1))
	}

	N := 99

	numer := new(big.Int)
	numer.Set(terms[N])
	denom := big.NewInt(1)

	new_numer := new(big.Int)
	new_denom := new(big.Int)

	for i := N - 1; i >= 0; i-- {
		log.Println(numer, " / ", denom)

		new_numer.Mul(terms[i], numer)
		new_numer.Add(new_numer, denom)

		// new_numer := terms[i] * numer + denom
		new_denom.Set(numer)

		_gcd := gcd(new_numer, new_denom)

		log.Println("new_numer: ", new_numer, " new_denom: ", new_denom, " _gcd: ", _gcd, " term: ", terms[i])

		numer.Div(new_numer, _gcd)
		denom.Div(new_denom, _gcd)
	}
	fmt.Println(numer, " / ", denom)
}
