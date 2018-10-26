package main

import (
	"time"
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Int())
	fmt.Println(rand.Int31())
	fmt.Println(rand.Int63())
	fmt.Println(rand.Float32())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Uint32())
	fmt.Println(rand.Uint64())
	fmt.Println(rand.Uint32())

	randomShuffle := strings.Fields("hello, world! gopher?")
	rand.Shuffle(len(randomShuffle), func(i, j int) {
		randomShuffle[i], randomShuffle[j] = randomShuffle[j], randomShuffle[i]
	})
	fmt.Println(randomShuffle)
}