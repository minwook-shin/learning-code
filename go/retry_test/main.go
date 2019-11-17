package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/kamilsk/retry"
	"github.com/kamilsk/retry/backoff"
	"github.com/kamilsk/retry/jitter"
	"github.com/kamilsk/retry/strategy"
)

func main() {
	action := func(uint) (err error) {
		defer func() {
			if r := recover(); r != nil {
				panic(r)
			}
		}()
		request, err := http.NewRequest(http.MethodGet, "http://example.com", nil)
		if err != nil {
			return err
		}
		response, err := http.DefaultClient.Do(request)
		fmt.Println(response)
		return err
	}

	strategies := retry.How{
		strategy.Limit(2),
		strategy.BackoffWithJitter(
			backoff.Fibonacci(5*time.Millisecond),
			jitter.NormalDistribution(
				rand.New(rand.NewSource(time.Now().UnixNano())), 0.25),
		),
		func(attempt uint, err error) bool {
			return attempt == 0 || err != nil
		},
	}

	breaker, _ := context.WithTimeout(context.Background(), time.Second)
	if err := retry.Try(breaker, action, strategies...); err != nil {
		panic(err)
	}
	fmt.Println("success")
}
