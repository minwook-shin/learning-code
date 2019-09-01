package main

import (
	"fmt"

	"github.com/Rhymond/go-money"
)

func main() {
	won := money.New(100, "KRW")
	doubleWon, err := won.Add(won)
	if err != nil {
		panic(err)
	}
	fmt.Println(doubleWon.Amount())
	fmt.Println(doubleWon.Currency().Code)
	fmt.Println(doubleWon.Currency().Formatter())
	fmt.Println(doubleWon.Currency().Grapheme)
	fmt.Println(doubleWon.Currency().Template)
	fmt.Println(doubleWon.Currency().Decimal)
	fmt.Println(doubleWon.Currency().Thousand)

	fmt.Println(doubleWon.Display())
	fmt.Println(doubleWon.IsPositive())

	fmt.Println(doubleWon.Negative().Display())
	fmt.Println(doubleWon.Negative().IsNegative())

	zero := money.New(0, "KRW")
	fmt.Println(zero.IsZero())
	fmt.Println(doubleWon.SameCurrency(zero))

	eq, err := doubleWon.Equals(doubleWon)
	if err != nil {
		panic(err)
	}
	fmt.Println(eq)

	gr, err := doubleWon.GreaterThan(won)
	if err != nil {
		panic(err)
	}
	fmt.Println(gr)

	less, err := won.LessThan(doubleWon)
	if err != nil {
		panic(err)
	}
	fmt.Println(less)

	sub, err := doubleWon.Subtract(won)
	if err != nil {
		panic(err)
	}
	fmt.Println(sub.Display())

	splitWon, err := doubleWon.Split(2)
	if err != nil {
		panic(err)
	}

	a := splitWon[0].Display()
	b := splitWon[1].Display()
	fmt.Println(a, b)
}
