package num7

import (
	"fmt"
	"time"
)

func num7() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("exit after %v\n", time.Since(start))

}

func or[T any](channels ...<-chan T) <-chan T {
	out := make(chan T)
	exitchan := make(chan struct{})

	output := func(c <-chan T) {
		for n := range c {
			out <- n
		}
		exitchan <- struct{}{}

	}
	for _, c := range channels {
		go output(c)
	}

	go func() {
		<-exitchan
		close(out)
	}()
	return out
}
