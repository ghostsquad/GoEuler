package main

//go:generate go build -buildmode=plugin -o main.so main.go

// The sum of the squares of the first ten natural numbers is,
//
// 1^2 + 2^2 + ... + 10^2 = 385
// The square of the sum of the first ten natural numbers is,
//
// (1 + 2 + ... + 10)^2 = 55^2 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 - 385 = 2640.
//
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

import (
	"context"
	"github.com/ghostsquad/goeuler/pkg"
)

type solution struct {}

// sum of squares formula
// 
// http://formulas.tutorvista.com/math/sum-of-squares-formula.html
// 
// 1^2 + 2^2 + 3^2 + .... + n^2 = ( n(n + 1)(2n + 1) ) / 6
// 
// sum of consective integers
// 
// (1 + 2 + 3 + ... + n)^2 = n(n+1)/2
// where N is the number of terms

func (s solution) Solve(ctx context.Context) {
	pkg.SolveWith(ctx, "006", func() uint64 {
		var answer uint64
		var n uint64 = 100
		var sumOfSquares uint64 = (n * (n + 1) * (2 * n + 1)) / 6;
		var sumOfConsecutive uint64 = (n * (n + 1)) / 2
		squareSum := sumOfConsecutive * sumOfConsecutive
		answer = squareSum - sumOfSquares

		return answer
	})
}

var Solution solution