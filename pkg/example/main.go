package main

//go:generate go build -buildmode=plugin -o main.so main.go

import (
	"fmt"
	"context"
	"github.com/ghostsquad/goeuler/pkg"
)

type solution struct {}

func (s solution) Solve(ctx context.Context) {
	pkg.SolveWith(ctx, "example", func() uint64 {
		var answer uint64
		var err error 
		if err != nil {
			fmt.Printf("An error occurred: %+v\n", err)
		}
		return answer
	})
}

var Solution solution