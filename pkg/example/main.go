package main

//go:generate go build -buildmode=plugin -o main.so main.go

import (
	"github.com/ghostsquad/goeuler/pkg"
)

type solution struct {}

func (s solution) Solve() {
	pkg.SolveWith("example", func() int {
		answer := 0
		return answer
	})
}

var Solution solution