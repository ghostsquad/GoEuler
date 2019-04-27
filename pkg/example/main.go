package main

//go:generate go build -buildmode=plugin -o main.so main.go

import (
	"context"
	"github.com/ghostsquad/goeuler/pkg"
)

type solution struct {}

func (s solution) Solve(ctx context.Context) {
	pkg.SolveWith(ctx, "example", func() uint64 {
		var answer uint64
		return answer
	})
}

var Solution solution