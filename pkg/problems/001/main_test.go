package main

import (
	"fmt"
	"testing"
)

// BenchmarkGetSumOfMultiples1-8   	233168
//  5000000	       244 ns/op
// BenchmarkGetSumOfMultiples2-8   	233168
//   100000	     15972 ns/op

func BenchmarkGetSumOfMultiples1(b *testing.B) {
	var answer int

	for n := 0; n < b.N; n++ {
		sumThrees := getSumOfMultiples1(3, 1000)
		sumFives := getSumOfMultiples1(5, 1000)
		sumDups := getSumOfMultiples1(15, 1000)
		
		answer = sumThrees + sumFives - sumDups
	}

	fmt.Println(answer)
}

func BenchmarkGetSumOfMultiples2(b *testing.B) {
	var answer int

	for n := 0; n < b.N; n++ {
		answer = getSumOfMultiples2(1000, 3, 5)
	}

	fmt.Println(answer)
}