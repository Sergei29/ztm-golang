package main

import (
	. "fmt"
)

func types() {
	type Price float32
	type Quantity uint32

	price := 22
	qty := 12.01
	convertedPrice := Price(price)
	convertedQty := Quantity(qty)
	Println(price, qty)
	Println(convertedPrice, convertedQty)
	sum := convertedPrice + 12.23
	Println("sum: ", sum)
}
