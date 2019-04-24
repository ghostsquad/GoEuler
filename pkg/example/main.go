package main

//go:generate go build -buildmode=plugin -o main.so main.go

import "fmt"

type solution struct {

}

func (s solution) Solve() {
    fmt.Println("Hello From Example")
}

var Solution solution