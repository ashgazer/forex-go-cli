package main

import (
	"flag"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var baseFlag = flag.String("b", "GBP", "your starting currency code")
	var destFlag = flag.String("d", "EUR", "your destination currency code")
	var conFlag = flag.String("c", "n", "do you wish to convert money?")
	flag.Parse()
	money, _ := strconv.ParseFloat(*conFlag, 64)


	var baseCurrerncy, DestCurrency string
	var baseCurrerncyP *string = &baseCurrerncy
	var DestCurrencyP *string = &DestCurrency

	*DestCurrencyP = strings.ToUpper(*destFlag)
	*baseCurrerncyP = strings.ToUpper(*baseFlag)

	rate := getRate(baseCurrerncy, DestCurrency)

	if int(money) > 0 {
	
		fmt.Printf("%.2f \n", calcChange(money, rate))

	}

}
