package main

//go:generate go build -buildmode=plugin -o main.so main.go

import (
	"fmt"
	"context"
	"github.com/ghostsquad/goeuler/pkg"
)

// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.

type solution struct {}

func (s solution) Solve(ctx context.Context) {
	pkg.SolveWith(ctx, "010", func() uint64 {
		var answer uint64

		p := pkg.NewPrimes()

		var twoMil uint64 = 1000 * 1000 * 2

		resultsChan, errChan := p.GenerateUpToIncluding(ctx, twoMil - 1)
		for range resultsChan {}

		if err := <-errChan; err != nil {
			fmt.Printf("received error from generate! err: %v", err)
			return 0
		}

		for _, prime := range p.Known {
			if prime >= twoMil {
				break
			}
			answer += prime
		}

		return answer
	})
}

var Solution solution