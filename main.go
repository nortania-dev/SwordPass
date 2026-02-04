package main

import (
	"fmt"
	"flag"
	"math"
	"math/rand"
)

func main() {
	p := flag.Float64("p", 1e-9, "acceptable probability of password compromise")
	v := flag.Float64("v", 100, "bruteforce speed (guesses per second)")
	t := flag.Float64("t", 7*24*3600, "password lifetime in seconds")
	a := flag.Int("a", 62, "alphabet size (26, 36, 62, 95)")

	flag.Parse()

	L := requiredLength(*v, *t, *p, *a)
	password := generatePassword(L, buildAlphabet(*a))

	fmt.Println(password)
}

func requiredLength(V, T, P float64, A int) int {
	S := (V * T) / P
	L := math.Log(S) / math.Log(float64(A))
	return int(math.Ceil(L))
}

func buildAlphabet(A int) []rune {
	switch A {
	case 26:
		return []rune("abcdefghijklmnopqrstuvwxyz")
	case 36:
		return []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	case 62:
		return []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	case 95:
		return []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ ")
	default:
		panic("unsupported alphabet size")
	}
}

func generatePassword(length int, alphabet []rune) string {
	pwd := make([]rune, length)
	for i := range pwd {
		pwd[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(pwd)
}
