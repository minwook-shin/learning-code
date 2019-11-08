package main

import (
	"fmt"
	"time"

	circuit "github.com/rubyist/circuitbreaker"
)

func main() {
	circuitBreaker := circuit.NewThresholdBreaker(1)

	BreakerEvent := circuitBreaker.Subscribe()
	go func() {
		for {
			<-BreakerEvent
		}
	}()

	for {
		if circuitBreaker.Ready() {
			httpClient := circuit.NewHTTPClient(time.Second*10, 10, nil)

			resp, err := httpClient.Get("http://example.com")
			fmt.Println(resp)

			if err != nil {
				circuitBreaker.Fail()
				continue
			}

			circuitBreaker.Success()
		}
	}
}
