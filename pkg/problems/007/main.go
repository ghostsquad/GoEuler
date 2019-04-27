package main

//go:generate go build -buildmode=plugin -o main.so main.go

// 10001st prime
// Problem 7
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

// What is the 10,001st prime number?

import (
	"fmt"
	"context"
	"github.com/ghostsquad/goeuler/pkg"
)

type solution struct {}

func (s solution) Solve(ctx context.Context) {
	pkg.SolveWith(ctx, "example", func() uint64 {
		var answer uint64
		var terms uint64 = 10001

		p := pkg.NewPrimes()

		primesChan, errChan := p.GenerateCountOf(ctx, terms)
		for range primesChan {}
		if err := <-errChan; err != nil {
			fmt.Printf("received error from generate! err: %v\n", err)
		}

		// don't get confused here
		// we are still getting correct term
		// but slices use 0-based indexes
		answer = p.Known[terms - 1]

		return answer
	})
}

var Solution solution