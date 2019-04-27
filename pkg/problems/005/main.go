package main

//go:generate go build -buildmode=plugin -o main.so main.go

// 2520 is the smallest number that can be divided by each of the numbers 
// from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible 
// by all of the numbers from 1 to 20?

import (
	"fmt"
	"math"
	"context"
	"github.com/ghostsquad/goeuler/pkg"
)

type solution struct {}


/*
* loop through each number, and factorize it
* example:
* 6=2*3
* 8=2*2*2 OR (2^3)
* 14 = 2*7
* 12 = 2*6 = 2*2*3 OR 2^2 * 3
*/

func (s solution) Solve(ctx context.Context) {
	pkg.SolveWith(ctx, "005", func() uint64 {
		var answer uint64
		var primeFactorial uint64

		factorPowersFinal := map[uint64]uint64{}

		p := pkg.NewPrimes()

		for i := uint64(2); i <= 20; i++ {
			factorPowers := map[uint64]uint64{}

			factorChan, errChan := p.Factorize(ctx, i)

			factors := []uint64{}
			for factor := range factorChan {
				factors = append(factors, factor)
				if _, ok := factorPowers[factor]; ok {
					factorPowers[factor]++
				} else {
					factorPowers[factor] = 1
				}
			}

			if err := <-errChan; err != nil {
				fmt.Printf("received error from generate! err: %v\n", err)
			}

			for factor, factorPower := range factorPowers {
				if val, ok := factorPowersFinal[factor]; !ok || val < factorPower {
					factorPowersFinal[factor] = factorPower
				}
			}
		}

		for factor, factorPower := range factorPowersFinal {
			primeFactorial = uint64(math.Pow(float64(factor), float64(factorPower)))

			if answer == 0 {
				answer = primeFactorial
			} else {
				answer = answer * primeFactorial
			}
		}

		// fmt.Printf("%+v\n", factorPowersFinal)

		return answer
	})
}

var Solution solution