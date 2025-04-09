package main

import (
	"fmt"
	"strings"
)
	const exchangeRateEUR = 0.9     
	const exchangeRateRUB = 85.0   
func main() {
	// amountUSD := 1.0
 

currency :=	getCurrency()
currentNumber :=	getNumber()
targetCurrent :=	getTargetCurrency(currency)

result := converter(currency, currentNumber, targetCurrent)

fmt.Printf("Результат: %.2f", result)
}

func getCurrency() string{
	var currency string
	for {
	fmt.Println("Введите исходную валюту (EUR/USD/RUB)")
	fmt.Scan(&currency)
	currency = strings.ToUpper(currency)	
	if currency == "EUR" || currency == "USD"  || currency == "RUB" {
		break
	}
	fmt.Println("Ошибка: некорректная валюта, попробуйте снова")
	}
	return currency
}

func getNumber() float64{
	var currentNumber float64

	for {
		fmt.Println("Введите количество валюты:")
		fmt.Scan(&currentNumber)

		if currentNumber > 0 {
			break
		}
		fmt.Println("Ошибка: некорректно передано число")
	}
	return currentNumber
}

func getTargetCurrency(currency string)(string)  {
	var targetCurrency string

	for {
		fmt.Println("Введите целевую валюту (EUR/USD/RUB):")
		fmt.Scan(&targetCurrency)
		targetCurrency = strings.ToUpper(targetCurrency)
		if targetCurrency == "EUR" || targetCurrency == "RUB" || targetCurrency == "USD" {
			if targetCurrency != currency {
				break
			}
		}
		fmt.Println("Ошибка: некорректно передана целевая валюта. Повторите ввод")
	}
	return targetCurrency
}

func converter(currency string, currentNumber float64, targetCurrent string) (float64) {
	var result float64
	if currency == "USD" {
		result = currentNumber
	}
	if currency == "EUR" {
		result = currentNumber / exchangeRateEUR
	}
	if currency == "RUB" {
		result = currentNumber / exchangeRateRUB
	}
	switch targetCurrent {
	case "EUR": 
		result = result * exchangeRateEUR
	case "RUB":
		result = result *exchangeRateRUB
	default:
		return result
	}
	return result
}