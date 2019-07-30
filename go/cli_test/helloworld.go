package main

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper

	Str    string         `cli:"*str" usage:"print string" dft:"hello"`
	I      int            `cli:"*i,integer" usage:"print integer" dft:"1"`
	B      bool           `cli:"*b" usage:"print boolean" dft:"True"`
	StrArr []string       `cli:"s,strArr" usage:"print string array`
	M      map[string]int `cli:"m" usage:"print map"`

	Version bool `cli:"!v" usage:"note the version"`
}

func (argv *argT) Validate(ctx *cli.Context) error {
	if argv.I > 10 {
		return fmt.Errorf("out of range")
	}
	if argv.B == false {
		return fmt.Errorf("invalid")
	}
	return nil
}

func main() {
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		if argv.Version {
			ctx.String("1.0\n")
			return nil
		}
		ctx.JSONln(ctx.Argv())
		return nil
	}))
}
