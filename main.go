package main

import (
	"fmt"
	"CurrencyConvert/lib"
)

type ConvertedCurrency struct {
	NewAmount float64 `json:"new_amount"`
	NewCurrency string `json:"new_currency"`
	OldCurrency string `json:"old_currency"`
	OldAmount float64 `json:"old_amount"`
}

func main() {
	var userName string

	fmt.Printf("Please enter your username: ")
	fmt.Scan(&userName)
	fmt.Printf("\nHey %v,\nWelcome to Currency Exchange,\n", userName)
	fmt.Printf("We charge 2%% on conversions below 100 USD and 1%% on above.\n\n")

	var oldCurrencyCode string
	var oldCurrency float64
	var newCurrencyCode string

	for {
		fmt.Printf("Enter the currency code of the currency you have: ")
		fmt.Scan(&oldCurrencyCode)
		fmt.Printf("Enter the value of currency you want to exchange: ")
		fmt.Scan(&oldCurrency)
		fmt.Printf("Enter the currency code of the currency you need: ")
		fmt.Scan(&newCurrencyCode)

		isOldCurrencyCodeValid, isOldCurrencyValid, isNewCurrencyCodeValid := lib.ValidateUserInput(oldCurrencyCode, oldCurrency, newCurrencyCode)

		if isOldCurrencyCodeValid && isOldCurrencyValid && isNewCurrencyCodeValid {
			var convertedCurrency ConvertedCurrency
			var exchangeCost float64
			
			lib.ConvertCurrency(oldCurrencyCode, oldCurrency, newCurrencyCode, &convertedCurrency)

			if oldCurrency > 100 { exchangeCost = convertedCurrency.NewAmount * 0.02} else { exchangeCost = convertedCurrency.NewAmount * 0.01 }

			fmt.Printf("\nTotal Value:    %.2f\nExchange cost: -%.2f\n----------------------\nNet Amount:     %.2f\n----------------------\n", 
			convertedCurrency.NewAmount, exchangeCost, convertedCurrency.NewAmount - exchangeCost)
			fmt.Printf("\nThank you!")
			
			break
		} else {
			if !isOldCurrencyCodeValid {
				fmt.Println("Old currency code is invalid.")
			}

			if !isOldCurrencyValid {
				fmt.Println("Value of old currency should be grater than zero.")
			}

			if !isNewCurrencyCodeValid {
				fmt.Println("New currency code is invalid.")
			}

			fmt.Println("Please enter details again")
		}
	}
}