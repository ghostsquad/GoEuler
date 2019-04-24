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

func (s solution) Solve() {
    fmt.Println("Hello From 002")

    defer func(begin time.Time) {
        fmt.Printf("took: %s\n", tookSeconds(time.Since(begin)))
    }(time.Now())

    term1 := 1
    term2 := 2
    var tempTerm int
    fourMil := 4000000
    var answer int

    for term2 <= fourMil {
        if term2 % 2 == 0 {
            answer += term2
        }

        tempTerm = term1
        term1 = term2
        term2 = tempTerm + term2
    }

    fmt.Printf("%d\n", answer)
}

var Solution solution