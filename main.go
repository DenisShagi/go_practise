package main

import "fmt"

func main() {
	amountUSD := 1.0
	exchangeRateEUR := 0.9     
	exchangeRateRUB := 85.0    

	amountEUR := amountUSD * exchangeRateEUR
	amountRUB := amountUSD * exchangeRateRUB
	amountEURinRUB := amountEUR / exchangeRateEUR * amountRUB

	fmt.Printf("EUR in RUB: %.2f\n", amountEURinRUB)
}
