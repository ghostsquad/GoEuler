package main

//go:generate go build -buildmode=plugin -o main.so main.go

import (
	"strconv"
	"context"
	"github.com/ghostsquad/goeuler/pkg"
)

// This solution decrements only 1 factor at a time, until a palindrome is found.
// Then it decrements both from the starting position, and continues.
// If a palindrome is found at anytime and
// a product lower than the lowest product from the previos run is created
// we exit, as
// it's not possible to create a larger palindrome once we reached that threshold

// TODO implement logging

type solution struct {}

func (s solution) Solve(ctx context.Context) {
	pkg.SolveWith(ctx, "004", func() uint64 {
		var answer uint64
		var factor1StartPoint uint64 = 999
		factor2StartPoint := factor1StartPoint
		var biggestProductPalindromeFromPreviousRun uint64

		thresholdReached := false

		for factor1 := factor1StartPoint; factor1 > 0; factor1-- {
			factor2StartPoint = factor1

			var product uint64

			for factor2 := factor2StartPoint; factor2 > 0; factor2-- {
				//fmt.Printf("Checking: %d x %d\n", factor1, factor2)

				product = factor1 * factor2
                //products are only getting smaller for each iteration
                //break early if the product is smaller than the biggest product from the previous run
				if product < biggestProductPalindromeFromPreviousRun {
					//fmt.Printf("Lowest product previous run has been surpassed: %d < %d\n", product, biggestProductPalindromeFromPreviousRun)

					// the threshold is reached if we've started this loop, and we are producing products
					// lower then lowest produced products from previous runs
					if factor2 == factor2StartPoint {
						//fmt.Printf("Reached threshold: %d x %d = %d\n", factor1, factor2, product)
						thresholdReached = true
					}

					break
				}

				if isPalindrome(product) {
					//fmt.Printf("Found palindrome: %d\n", product)

					if product > answer {
						biggestProductPalindromeFromPreviousRun = product
						answer = product
					}
					
					break
				}
			}

			if thresholdReached {
				break
			}
		}

		return answer
	})
}

func isPalindrome(x uint64) bool {
	str := strconv.FormatUint(x, 10)
	rev := reverse(str)

	return str == rev
}

func reverse(s string) string {
    n := len(s)
    runes := make([]rune, n)
    for _, rune := range s {
        n--
        runes[n] = rune
    }
    return string(runes[n:])
}

var Solution solution