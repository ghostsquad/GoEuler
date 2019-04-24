package pkg

import (
	"fmt"
	"time"
)

func tookSeconds(d time.Duration) string {
	return fmt.Sprintf("%f seconds", d.Seconds())
}

func SolveWith(name string, solver func() int) {
	fmt.Printf("Hello From %s\n", name)
	defer func(begin time.Time) {
        fmt.Printf("took: %s\n", tookSeconds(time.Since(begin)))
	}(time.Now())

	answer := solver()
	
	fmt.Printf("%d\n", answer)
}
