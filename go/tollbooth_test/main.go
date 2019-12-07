package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
)

func handler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	newLimiter := tollbooth.NewLimiter(1, nil)
	newLimiter = tollbooth.NewLimiter(1, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})

	newLimiter.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
	newLimiter.SetMethods([]string{"GET", "POST"})

	newLimiter.SetTokenBucketExpirationTTL(time.Hour)
	newLimiter.SetBasicAuthExpirationTTL(time.Hour)
	newLimiter.SetHeaderEntryExpirationTTL(time.Hour)

	newLimiter.SetMessage("limit is reached")
	newLimiter.SetMessageContentType("text/plain; charset=utf-8")

	newLimiter.SetOnLimitReached(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Set On Limit Reached")
	})

	http.Handle("/", tollbooth.LimitFuncHandler(newLimiter, handler))

	http.ListenAndServe(":8000", nil)
}
