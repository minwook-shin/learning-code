package main

import (
	"fmt"
	"time"

	"github.com/caarlos0/env"
)

type config struct {
	String   string        `env:"HOME"`
	Int      int           `env:"INT" envDefault:"1"`
	Bool     bool          `env:"PRODUCTION" envDefault:"true"`
	Slice    []string      `env:"SLICE" envSeparator:":" envDefault:"hello:world"`
	Duration time.Duration `env:"DURATION"`
}

func main() {
	config := config{}
	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", config)
}
