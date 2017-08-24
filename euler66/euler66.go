package main

import "fmt"
import "math/big"
import "log"
import "container/list"
import "bytes"

const D_LIMIT = 1001

var small []big.Int
var smallSquare []big.Int
var sqrt [D_LIMIT]int

func GetInt(x int) *big.Int {
	L := len(small)
	if x >= L {
		for i := L; i < L * 2; i++ {
			small = append(small, *big.NewInt(int64(i)))
		}
	}

	return &small[x]
}

func GetSquare(x int) *big.Int {
	L := len(smallSquare)
	if x >= L {
		for i := L; i < L * 2; i++ {
			var square big.Int
			I := GetInt(i)

			square.Mul(I, I)
			smallSquare = append(smallSquare, square)
		}
	}

	return &smallSquare[x]
}

func is_square(N int) bool {
	i := 0
	for ; i * i < N; i++ {
	}
	return i * i == N
}

func init() {
	for i := 0; i < 10; i++ {
		small = append(small, *big.NewInt(int64(i)))
		// _ = GetInt(i)
	}

	for i := 0; i < 10; i++ {
		var S big.Int
		S.Mul(GetInt(i), GetInt(i))
		smallSquare = append(smallSquare, S)
		// _ = GetSquare(i)
	}

	for i := 0; i < D_LIMIT; i++ {
		j := 0
		for ; j * j <= i; j++ {
		}
		if j * j == i {
			sqrt[i] = j
		} else {
			sqrt[i] = j - 1
		}
	}
}

func new_list_of_non_squares(begin int, end int) *list.List {
	result := list.New()
	for i := begin; i < end; i++ {
		if is_square(i) {
			continue
		}
		result.PushBack(i)
	}
	return result
}

func List2String(self list.List) string {
	buff := new(bytes.Buffer)
	for node := self.Front(); node != nil; node = node.Next() {
		fmt.Fprint(buff, node.Value)
		fmt.Fprint(buff, " ")
	}
	return buff.String()
}

func x_solves_D(x int64, D int64) bool {
	// x solves D if for some integer y, y**2 == (x**2 - 1) / D
	target := x * x - 1
	if target % D != 0 {
		return false
	}
	target /= D

	// Easy to prove y has an upper limit of x / sqrt[D]
	y := x / int64(sqrt[D]) + 1
	tries := 0
	for ; y * y > target; y-- {
		tries += 1
	}

	if y * y == target {
		log.Printf("%d^2 - %dx%d^2 = 1\n", x, D, y)
	}

	if x % 100000 == 0 {
		log.Println ("x_solves_D", tries, "tries")
	}

	return y * y == target
}

func main() {
	D_list := new_list_of_non_squares(2, 1001)

	for x := int64(2); D_list.Len() > 1; x++ {
		if x % 100000 == 0 {
			log.Println ("x=",x)
			log.Println("[", List2String(*D_list), "]")
		}
		var next_node *list.Element
		for node := D_list.Front(); node != nil; node = next_node {
			next_node = node.Next()

			if x_solves_D(x, int64(node.Value.(int))) {
				D_list.Remove(node)
			}
		}

		if D_list.Len() == 1 {
			break
		}
	}

	fmt.Println(D_list.Front().Value)
}



func minimal(D int) big.Int {
	var total big.Int

	var bigD big.Int
	bigD.SetInt64(int64(D))

	x := 1
	for y := 1; ; y++ {
		//log.Println("Y: ", y)
		total.Mul(&smallSquare[y], &bigD)
		total.Add(&total, &small[1])

		for {
			// log.Println("X: ", x)
			_cmp := GetSquare(x).Cmp(&total)
			if _cmp == 0 {
				return *GetInt(x)
			}
			if _cmp > 0 {
				break
			}
			x++

		}
	}

	log.Fatal ("Couldn't find a solution for D =", D)

	return small[0]
}
