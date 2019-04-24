package main

//go:generate go build -buildmode=plugin -o main.so main.go

import (
	"context"
	"time"
	"github.com/ghostsquad/goeuler/pkg"
)

type solution struct {}

func (s solution) Solve(ctx context.Context) {
	pkg.SolveWith(ctx, "003", func() int {
		answer := 0

		time.Sleep(5 * time.Second)

		return answer
	})
}

var Solution solution