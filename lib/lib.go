package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)                     

func ValidateUserInput (oldCurrencyCode string, oldCurrency float64, newCurrencyCode string) (bool, bool, bool) {
	isOldCurrencyCodeValid, _ :=  regexp.MatchString("^[A-Z]{3}$", oldCurrencyCode) 
	isOldCurrencyValid :=  oldCurrency > 0
	isNewCurrencyCodeValid, _ :=  regexp.MatchString("^[A-Z]{3}$", newCurrencyCode)
	
	return isOldCurrencyCodeValid, isOldCurrencyValid, isNewCurrencyCodeValid
}

func ConvertCurrency (oldCurrencyCode string, oldCurrency float64, newCurrencyCode string,target interface{} ) error {
	url := fmt.Sprintf("https://currency-converter-by-api-ninjas.p.rapidapi.com/v1/convertcurrency?have=%v&want=%v&amount=%v",
	oldCurrencyCode, newCurrencyCode, oldCurrency)

	req, error := http.NewRequest("GET", url, nil)

	if error != nil {
		fmt.Println(error)
	}

	req.Header.Add("X-RapidAPI-Key", "ed0601ffa0msh56198aa82db94cbp1d6622jsnb87d159e3561")
	req.Header.Add("X-RapidAPI-Host", "currency-converter-by-api-ninjas.p.rapidapi.com")

	res, error := http.DefaultClient.Do(req)

	if error != nil {
		fmt.Println(error)
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}