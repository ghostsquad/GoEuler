package main

//go:generate go build -buildmode=plugin -o main.so main.go

import (
    "context"
    "github.com/ghostsquad/goeuler/pkg"
)

type solution struct {}

func (s solution) Solve(ctx context.Context) {
    pkg.SolveWith(ctx, "002", func() int {
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

        return answer
    })
}

var Solution solution