package main

import (
	crypto "crypto/rand"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Intn(100))

	fmt.Println(rand.Float64())

	fmt.Println((rand.Float64() * 5) + 5)

	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	fmt.Println(random.Intn(100))

	cryptoValue := make([]byte, 1)
	crypto.Read(cryptoValue)
	fmt.Println(cryptoValue[0])

	firstSource := rand.NewSource(2018)
	firstRandom := rand.New(firstSource)
	fmt.Println(firstRandom.Intn(100))

	secondSource := rand.NewSource(2018)
	secondRandom := rand.New(secondSource)
	fmt.Println(secondRandom.Intn(100))
}
