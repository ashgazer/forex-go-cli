package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/ashgazer/forexgocli/utils"
)

func main() {
	var baseFlag = flag.String("b", "GBP", "your starting currency code")
	var destFlag = flag.String("d", "GBP", "your destination currency code")
	var conFlag = flag.String("c", "0.00", "the amount you wish to convert")
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

	rate := utils.GetRate(baseCurrerncy, DestCurrency)

	if int(money) > 0 {

		fmt.Printf("You will get %.2f %s\n", utils.CalcChange(money, rate), DestCurrency)

	}

}
