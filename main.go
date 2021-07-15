package main

import (
	"rsc.io/quote"
	"tenyunops.com/nashyang/calculator"
)

func main() {
	println(calculator.Sum(3, 5))
	println("Version: ", calculator.Version)
	println(quote.Hello())
}
