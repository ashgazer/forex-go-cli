package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var baseFlag = flag.String("b", "GBP", "your starting currency code")
	var destFlag = flag.String("d", "GBP", "your destination currency code")
	var conFlag = flag.String("c", "n", "do you wish to convert money?")
	flag.Parse()
	
	if *destFlag == *baseFlag {
		os.Exit(0)
	}
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
