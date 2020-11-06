package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var baseFlag = flag.String("b", "GBP", "your starting currency code")
	var destFlag = flag.String("d", "EUR", "your destination currency code")
	var conFlag = flag.String("c", "n", "do you wish to convert money?")
	flag.Parse()

	var baseCurrerncy, DestCurrency string
	var baseCurrerncyP *string = &baseCurrerncy
	var DestCurrencyP *string = &DestCurrency

	*DestCurrencyP = strings.ToUpper(*destFlag)
	*baseCurrerncyP = strings.ToUpper(*baseFlag)

	rate := getRate(baseCurrerncy, DestCurrency)

	if *conFlag == "y" {
		var money float64
		fmt.Println("How much money do you want to change?")

		fmt.Scanln(&money)
		fmt.Printf("%.2f \n", calcChange(money, rate))

	}

}
