package main

import (
	"context"
	"fmt"
	"os"
	"plugin"
	"strconv"
	"github.com/spf13/cobra"
)

func NewSolveCmd(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "solve PROBLEM",
		Short: "Solve a problem",
		Run: func(cmd *cobra.Command, args []string) {
			problemNumStr := args[0]
	
			problemNum, err := strconv.Atoi(problemNumStr)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
	
			// load module
			// 1. open the so file to load the symbols
			problemFileNoExt := fmt.Sprintf("%03d", problemNum)
			plug, err := plugin.Open(fmt.Sprintf("./pkg/problems/%s/main.so",  problemFileNoExt))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
	
			symSolution, err := plug.Lookup("Solution")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
	
			var solution Solution
			solution, ok := symSolution.(Solution)
			if !ok {
				fmt.Fprintln(os.Stderr, "unexpected type from module symbol")
				os.Exit(1)
			}
	
			solution.Solve(ctx)
		},
	}
}
