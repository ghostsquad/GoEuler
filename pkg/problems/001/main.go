package main

//go:generate go build -buildmode=plugin -o main.so main.go

import "fmt"

type solution struct {

}

func getSumOfMultiples(multipleOf, max int) (sum int) {
	for i := multipleOf; i < max; i += multipleOf {
		sum += i
	}

	return
}

func (s solution) Solve() {
	fmt.Println("Hello From 001")

	sumThrees := getSumOfMultiples(3, 1000)
	sumFives := getSumOfMultiples(5, 1000)
	sumDups := getSumOfMultiples(15, 1000)
	
	answer := sumThrees + sumFives - sumDups

	fmt.Printf("%d\n", answer)
}

var Solution solution