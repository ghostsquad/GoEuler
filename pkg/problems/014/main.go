package main

//go:generate go build -buildmode=plugin -o main.so main.go

import (
	"context"
	"fmt"
	"github.com/ghostsquad/goeuler/pkg"
)

type solution struct{}

//The following iterative sequence is defined for the set of positive integers:

//n → n/2 (n is even)
//n → 3n + 1 (n is odd)

//Using the rule above and starting with 13, we generate the following sequence:

//13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
//It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

//Which starting number, under one million, produces the longest chain?

//NOTE: Once the chain starts the terms are allowed to go above one million.

func (s solution) Solve(ctx context.Context) {
	pkg.SolveWith(ctx, "014", func() uint64 {
		var answer uint64
		var longestChainLength uint64

		for startingNum := uint64(1000 * 1000); startingNum > 0; startingNum-- {
			next := startingNum
			chainLength := uint64(0)

			for next > 1 {
				chainLength++
				next = nextNumberInSequence(next)
			}

			if chainLength > longestChainLength {
				longestChainLength = chainLength
				answer = startingNum
			}
		}

		var err error
		if err != nil {
			fmt.Printf("An error occurred: %+v\n", err)
		}
		return answer
	})
}

func nextNumberInSequence(current uint64) (result uint64) {
	if current%2 == 0 {
		result = current / 2
	} else {
		result = (3 * current) + 1
	}

	return
}

var Solution solution
