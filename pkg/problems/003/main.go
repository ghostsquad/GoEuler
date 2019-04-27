package main

//go:generate go build -buildmode=plugin -o main.so main.go

// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?

import (
	"fmt"
	"math"
	"context"
	"github.com/ghostsquad/goeuler/pkg"
)

type solution struct {}

func (s solution) Solve(ctx context.Context) {
	pkg.SolveWith(ctx, "003", func() int {
		var answer int
		var num uint64 = 600851475143
		sqrt := uint64(math.Sqrt(float64(num)))

		kp := pkg.NewPrimes()
		
		resultsChan, errChan := kp.GenerateUpToIncluding(ctx, sqrt)
		for range resultsChan {}
		if err := <-errChan; err != nil {
			fmt.Printf("received error from generate! err: %v\n", err)
			return 0
		}

		startingI := len(kp.Known) - 1
		for i := startingI; i > 0; i-- {
			if kp.Known[i] <= sqrt {
				startingI = i
				break
			}
		}

		for i := startingI; i > 0; i-- {
			prime := kp.Known[i]
			if num % prime == 0 {
				answer = int(prime)
				break
			}
		}

		return answer
	})
}

var Solution solution