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
    fmt.Println("Hello From Example")

    defer func(begin time.Time) {
        fmt.Printf("took: %s\n", tookSeconds(time.Since(begin)))
    }(time.Now())
}

var Solution solution