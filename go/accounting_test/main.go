package main

import (
	"fmt"

	"github.com/leekchan/accounting"
)

func main() {
	AccountingStruct := accounting.Accounting{Symbol: "₩"}
	fmt.Println(AccountingStruct.FormatMoney(123456789))

	AccountingStruct = accounting.Accounting{Symbol: "₩", Precision: 0, Thousand: ".", Decimal: ","}
	fmt.Println(AccountingStruct.FormatMoney(123456789))

	AccountingStruct = accounting.Accounting{Symbol: "KRW", Format: "%s %v", FormatNegative: "%s -%v", FormatZero: "%s --"}
	fmt.Println(AccountingStruct.FormatMoney(123456789))
	fmt.Println(AccountingStruct.FormatMoney(-1000))
	fmt.Println(AccountingStruct.FormatMoney(0))

	DefaultAccounting := accounting.DefaultAccounting("KRW", 0)
	fmt.Println(DefaultAccounting.FormatMoney(123456789))
	fmt.Println(DefaultAccounting.FormatMoney(-1000))
	fmt.Println(DefaultAccounting.FormatMoney(0))

	DefaultAccounting = accounting.NewAccounting("KRW", 0, ",", ".", "%s %v", "%s -%v", "%s --")
	fmt.Println(DefaultAccounting.FormatMoney(123456789))
	fmt.Println(DefaultAccounting.FormatMoney(-1000))
	fmt.Println(DefaultAccounting.FormatMoney(0))
}
