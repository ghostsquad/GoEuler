package main

//go:generate go build -buildmode=plugin -o 001.so 001.go

import "fmt"

type solution struct {

}

func (s solution) Solve() {
    fmt.Println("Hello From 001")
}

var Solution solution