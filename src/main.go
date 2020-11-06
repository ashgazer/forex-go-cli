package main

import "fmt"
import "strings"

func main() {

	var baseCurrerncy, DestCurrency, ans string
	var baseCurrerncyP *string = &baseCurrerncy
	var DestCurrencyP *string = &DestCurrency
	fmt.Println("what is the currency you want to change from?")
	fmt.Scanln(&baseCurrerncy)
	fmt.Println("what is the currency you want to change to?")
	fmt.Scanln(&DestCurrency)

	*DestCurrencyP = strings.ToUpper(DestCurrency)
	*baseCurrerncyP = strings.ToUpper(baseCurrerncy)

	rate := getRate(baseCurrerncy, DestCurrency)

	fmt.Println("\n\nDo you want to convert your money? y/n")
	fmt.Scanln(&ans)

	if ans == "y" {
		var money float64
		fmt.Println("How much money do you want to change?")

		fmt.Scanln(&money)
		fmt.Printf("%.2f \n", calcChange(money, rate))

	}

}
