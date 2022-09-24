package modules

import "strconv"

var EfisheryGoods = []map[string]int{
	{"Benih Lele": 50000},
	{"Pakan Lele Cap Menara": 25000},
	{"Probiotik": 75000},
	{"Probiotik Nila B": 10000},
	{"Pakan Nila": 20000},
	{"Benih Nila": 20000},
	{"Cupang": 5000},
	{"Benih Nila": 30000},
	{"Benih Cupang": 10000},
	{"Probiotik B": 10000},
}

// Function A
func GoodsFromMoney(moneyUser int) []string {
	var result []string // array for result
	for _, goods := range EfisheryGoods {
		for key, price := range goods {
			if moneyUser >= price {
				moneyUser = moneyUser - price
				result = append(
					result,
					key,
					"-",
					strconv.Itoa(price),
				)
			} else {
				continue
			}
		}
	}
	return result
}

// Function B
// return lowest price and highest price
func LowHigh() []string {
	var lowestPrice int
	var highestPrice int
	var lowestGoods string
	var highestGoods string

	for _, goods := range EfisheryGoods {
		for key, price := range goods {
			if lowestPrice == 0 {
				lowestPrice = price
				lowestGoods = key
			}
			if highestPrice == 0 {
				highestPrice = price
				highestGoods = key
			}
			if lowestPrice > price {
				lowestPrice = price
				lowestGoods = key
			}
			if highestPrice < price {
				highestPrice = price
				highestGoods = key
			}
		}
	}
	var result = []string{
		lowestGoods + "- " + strconv.Itoa(lowestPrice),
		highestGoods + "- " + strconv.Itoa(highestPrice),
	}
	return result
}

// Function C
// Getting key from EfisheryGoods that have 10000 value using for loop output array
func GetGoodsFromPrice(value int) []string {
	var result []string
	for _, goods := range EfisheryGoods {
		for key, price := range goods {
			if price == value {
				result = append(result, key, "-", strconv.Itoa(price))
			}
		}
	}
	return result
}
