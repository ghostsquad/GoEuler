package main

//go:generate go build -buildmode=plugin -o main.so main.go

import (
	"math"
	"context"
	"github.com/ghostsquad/goeuler/pkg"
)

// Special Pythagorean triplet
// Problem 9
// A Pythagorean triplet is a set of three natural numbers, a  b  c, for which,

// a2 + b2 = c2
// For example, 32 + 42 = 9 + 16 = 25 = 52.

// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

// r = 150
// s = 50
// t = 225
// 
// a = 200
// b = 375
// c = 425
// 
// numToFactor 11250
// factors [50,225] = 11250
// triple generated [200,375,425]

// Pythagorean Theorem
// a^2 + b^2 = c^2
// a < b < c
// 
// triplet generation using Dickson method
// 
// http://en.wikipedia.org/wiki/Formulas_for_generating_Pythagorean_triples
// 
// r^2 = 2st
// 
// a = r + s
// b = r + t
// c = r + s + t
// 
// r is an EVEN integer
// 
// s & t are factors of r^2 / 2
// 
// Example: r = 6
// then r^2 / 2 = 18
// 
// factor pairs of 18 are (1,18), (2,9), and (3,6) all of which produce triples
// 
// s = 1, t = 18 produces the triple [7, 24, 25] because a = 6 + 1 = 7,  b = 6 + 18 = 24,  c = 6 + 1 + 18 = 25.
// s = 2, t = 9 produces the triple [8, 15, 17] because a = 6 + 2 = 8,  b = 6 +  9 = 15,  c = 6 + 2 + 9 = 17.
// s = 3, t = 6 produces the triple [9, 12, 15] because a = 6 + 3 = 9,  b = 6 +  6 = 12,  c = 6 + 3 + 6 = 15.

type solution struct {}

func (s solution) Solve(ctx context.Context) {
	pkg.SolveWith(ctx, "009", func() uint64 {
		var answer uint64

		sumGoal := 1000
		r := 2
		var s int
		var t int

		var a int
		var b int
		var c int

		var sum int
		var numToFactor int

		for answer == 0 {
			numToFactor = r * r / 2
			factors := generateFactors(numToFactor)
			
			for _, tuple := range factors {
				s = tuple.s
				t = tuple.t

				a = r + s
				b = r + t
				c = r + s + t

				sum = a + b + c
				if sum == sumGoal {
					answer = uint64(a * b * c)
					break
				}
			}

			r += 2
		}

		return answer
	})
}

type factorTuple struct {
	s int
	t int
}

func generateFactors(num int) []factorTuple  {
	tuples := []factorTuple{
		factorTuple{1, num},
	}

	sqrt := math.Sqrt(float64(num))
	numFloat := float64(num)

	for i := float64(2); i < sqrt; i++ {
		quotient := numFloat / i

		if math.Mod(quotient, 1) == 0 {
			tuples = append(tuples, factorTuple{int(i), int(quotient)})
		}
	}

	return tuples
}

var Solution solution