package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := errors.New("ERROR")
	fmt.Printf("%+v", err)

	fmt.Println("\n1**********")

	cause := errors.New("ERROR")
	err = errors.WithMessage(cause, "new message.")
	fmt.Printf("%+v", err)

	fmt.Println("\n2**********")

	cause = errors.New("ERROR")
	err = errors.WithStack(cause)
	fmt.Printf("%+v", err)

	fmt.Println("\n3**********")

	cause = errors.New("ERROR")
	err = errors.Wrap(cause, "wrap")
	fmt.Printf("%+v", err)

	fmt.Println("\n4**********")

	cause = errors.New("ERROR")
	err = errors.Wrapf(cause, "wrap : %d", 1)
	fmt.Printf("%+v", err)

	fmt.Println("\n5**********")

	err = errors.Errorf("ERROR : %s", "test")
	fmt.Printf("%+v", err)

	fmt.Println("\n6**********")

	err = errors.Wrap(func() error {
		return func() error {
			return errors.Errorf("ERROR : %s", fmt.Sprintf("test"))
		}()
	}(), "wrap")

	fmt.Printf("%+v", err)
}
