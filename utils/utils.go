package utils


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
)

// Rate struct
type Rate struct {
	Base  string `json:"base"`
	Rates struct {
		GBP float64 `json:"GBP"`
	} `json:"rates"`
	CDate string `json:"date"`
}

func round(value float64, places float64) float64 {
	var numOfPlaces float64 = math.Pow(10, places)
	cal := math.Round(value*numOfPlaces) / numOfPlaces
	return cal
}

func getXURL(baseCurrency string, destCurrency string) string {
	apiEndpoint := "https://api.ratesapi.io/api/latest"

	base, err := url.Parse(apiEndpoint)

	if err != nil {

		fmt.Print(err)
	}

	params := url.Values{}
	params.Add("base", baseCurrency)
	params.Add("symbols", destCurrency)
	params.Add("rtype", "fpy")
	base.RawQuery = params.Encode()
	return base.String()
}

// GetRate gets rates between two currencies 
func GetRate(baseCurrency string, destCurrency string) float64 {
	resp, err := http.Get(getXURL(baseCurrency, destCurrency))

	if err != nil {

		fmt.Print(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// var rateStruct Rate
	var data map[string]interface{}
	json.Unmarshal(bodyBytes, &data)
	rates := data["rates"].(map[string]interface{})
	rate := rates[destCurrency].(float64)
	fmt.Printf("the conversion for %s to %s is %f. :-)\n", baseCurrency, destCurrency, rate)

	return rate
}

// CalcChange used to calculate the change in currency
func CalcChange(money float64, rate float64) float64 {
	return money * rate

}