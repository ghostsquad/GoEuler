package main

//go:generate go build -buildmode=plugin -o main.so main.go

import (
	"fmt"
	"time"
)

type solution struct {}

func tookSeconds(d time.Duration) string {
	return fmt.Sprintf("%f seconds", d.Seconds())
}

func getSumOfMultiples(multipleOf, max int) (sum int) {
	for i := multipleOf; i < max; i += multipleOf {
		sum += i
	}

	return
}

func (s solution) Solve() {
	fmt.Println("Hello From 001")

	defer func(begin time.Time) {
        fmt.Printf("took: %s\n", tookSeconds(time.Since(begin)))
    }(time.Now())

	sumThrees := getSumOfMultiples(3, 1000)
	sumFives := getSumOfMultiples(5, 1000)
	sumDups := getSumOfMultiples(15, 1000)
	
	answer := sumThrees + sumFives - sumDups

	fmt.Printf("%d\n", answer)
}

var Solution solution