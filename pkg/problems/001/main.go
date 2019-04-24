package main

//go:generate go build -buildmode=plugin -o main.so main.go

import (
	"github.com/ghostsquad/goeuler/pkg"
)

type solution struct {}

func getSumOfMultiples(multipleOf, max int) (sum int) {
	for i := multipleOf; i < max; i += multipleOf {
		sum += i
	}

	return
}

func (s solution) Solve() {
	pkg.SolveWith("001", func() int {
		sumThrees := getSumOfMultiples(3, 1000)
		sumFives := getSumOfMultiples(5, 1000)
		sumDups := getSumOfMultiples(15, 1000)
		
		answer := sumThrees + sumFives - sumDups
		return answer
	})
}

var Solution solution