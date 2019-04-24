package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

func NewGenPrimesCmd(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "gen-primes",
		Short: "Generate Prime Numbers. Exit with Ctrl-C",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside gen-primes Run with args: %v\n", args)
		},
	}
}
