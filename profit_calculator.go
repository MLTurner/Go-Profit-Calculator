package main

import (
	"fmt"
)

func main() {

	var taxRate float64

	var revenue float64

	var expenses float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := (revenue - expenses)

	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)

	ratio := (earningsBeforeTax / earningsAfterTax)

	fmt.Println(earningsBeforeTax)
	fmt.Println(earningsAfterTax)
	fmt.Println(ratio)

}
