package main

import (
	"fmt"
	"os"
	"plugin"
	"strconv"
	"github.com/spf13/cobra"
)

type Solution interface {
	Solve()
}

var rootCmd = &cobra.Command{
	Use:   "goeuler",
	Short: "Run euler solutions",
	Run: func(cmd *cobra.Command, args []string) {
		problemNumStr := args[0]

		problemNum, err := strconv.Atoi(problemNumStr)
		if err != nil {
			fmt.Println(err)
		}

		// load module
		// 1. open the so file to load the symbols
		problemFileNoExt := fmt.Sprintf("%03d", problemNum)
		plug, err := plugin.Open(fmt.Sprintf("./pkg/problems/%s.so",  problemFileNoExt))
		if err != nil {
			fmt.Println(err)
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
			fmt.Println("unexpected type from module symbol")
			os.Exit(1)
		}

		solution.Solve()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
