package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func main() {
	localTime := time.Unix(1000000, 0)
	unixTime := localTime.UnixNano()

	randSource := rand.NewSource(unixTime)
	random := rand.New(randSource)

	entropy := ulid.Monotonic(random, 0)

	ulid := ulid.MustNew(ulid.Timestamp(localTime), entropy)
	fmt.Println(ulid)
}
