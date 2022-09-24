package main

import (
	"assignment/modules"
	"fmt"
)

func main() {
	//Assignment 1
	// Answer A
	fmt.Println(modules.GoodsFromMoney(
		100000,
	))
	// Answer B
	fmt.Println(modules.LowHigh())

	// Answer C
	fmt.Println(modules.GetGoodsFromPrice(
		10000, // value price
	))
}
