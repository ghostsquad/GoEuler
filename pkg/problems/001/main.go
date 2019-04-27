package main

//go:generate go build -buildmode=plugin -o main.so main.go

import (
	"context"
	"github.com/ghostsquad/goeuler/pkg"
)

type solution struct {}

func getSumOfMultiples1(multipleOf, max int) (sum int) {
	for i := multipleOf; i < max; i += multipleOf {
		sum += i
	}

	return
}

func getSumOfMultiples2(max int, multiples ...int) (sum int) {
	for i := 0; i < max; i ++ {
			for _, multiple := range multiples {
					if i % multiple == 0 {
							sum += i
							break
					}
			}
	}
	return
}

func (s solution) Solve(ctx context.Context) {
	pkg.SolveWith(ctx, "001", func() uint64 {
		var answer uint64

		sumThrees := getSumOfMultiples1(3, 1000)
		sumFives := getSumOfMultiples1(5, 1000)
		sumDups := getSumOfMultiples1(15, 1000)
		
		answer = uint64(sumThrees) + uint64(sumFives) - uint64(sumDups)
		return answer
	})
}

var Solution solution