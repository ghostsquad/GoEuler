package pkg

import (
	"os"
	"context"
	"fmt"
	"time"
)

func tookSeconds(d time.Duration) string {
	return fmt.Sprintf("%f seconds", d.Seconds())
}

func SolveWith(ctx context.Context, name string, solver func() uint64) {
	fmt.Fprintf(os.Stderr, "Hello From %s\n", name)
	defer func(begin time.Time) {
        fmt.Fprintf(os.Stderr, "took: %s\n", tookSeconds(time.Since(begin)))
	}(time.Now())

	answerChannel := make(chan uint64, 1)
	go func() {
        answerChannel <- solver()
	}()
	
	select {
    case answer := <-answerChannel:
		fmt.Fprint(os.Stderr, "answer: ")
		fmt.Fprintf(os.Stdout, "%d", answer)
		fmt.Fprint(os.Stderr, "\n")
    case <-ctx.Done():
        fmt.Fprint(os.Stderr, "Cancellation requested\n")
    }
}
