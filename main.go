package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/spf13/cobra"
	// a universal mechanism to manage goroutine lifecycles
	"github.com/oklog/run"
	"github.com/pkg/errors"
	"github.com/fatih/color"
)

type Solution interface {
	Solve(context.Context)
}

func newRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "goeuler [sub]",
		Short: "Run Project Euler solutions",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	rootCmd := newRootCommand()

	genPrimesCmd := NewGenPrimesCmd(ctx)
	solveCmd := NewSolveCmd(ctx)

	rootCmd.AddCommand(genPrimesCmd)
	rootCmd.AddCommand(solveCmd)

	runGroup := run.Group{}
	{
		cancelInterrupt := make(chan struct{})
		runGroup.Add(
			createSignalWatcher(ctx, cancelInterrupt, cancel), 
			func(error) {
				close(cancelInterrupt)
			})
	}
	{
		runGroup.Add(func() error {
			return rootCmd.Execute()
		}, func(error) {
			cancel()
		})
	}

	err := runGroup.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "exit reason: %s\n", err)
		os.Exit(1)
	}

	color.New(color.FgGreen).Fprintln(os.Stderr, "Done!")
}

// This function just sits and waits for ctrl-C
func createSignalWatcher(ctx context.Context, cancelInterruptChan <-chan struct{}, cancel context.CancelFunc) func() error {
	return func() error {
		c := make(chan os.Signal, 1)

		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			err := errors.Errorf("received signal %s", sig)
			fmt.Fprintf(os.Stderr, "%s\n", err)
			signal.Stop(c)
			cancel()
			return err
		case <-ctx.Done():
			return nil
		case <-cancelInterruptChan:
			return nil
		}
	}
}
