package main

import (
	"fmt"
	"math"
)

func calculate(amount int, coins []int, arr [19]int) (int, [19]int) {
	if amount == 0 {
		return 0, arr
	}
	numberOfCoins := math.MaxInt16

	for i := 0; i < len(coins); i++ {

		if (amount - coins[i]) >= 0 {
			calcoins := 0
			if (arr[amount-coins[i]]) != 0 {
				calcoins = arr[amount-coins[i]]
			} else {

				calcoins, arr = calculate(amount-coins[i], coins, arr)
			}
			//arr[i] = calcoins
			//println(calcoins)
			if calcoins+1 < numberOfCoins {
				numberOfCoins = calcoins + 1
			}

		}

	}
	arr[amount] = numberOfCoins
	return numberOfCoins, arr
}

func main() {
	amount := 18
	coins := []int{7, 5, 1}
	var arr [19]int
	arr[0] = 0

	numberOfCoins, arr1 := calculate(amount, coins, arr)
	println(numberOfCoins)
	fmt.Println(arr1)
	//fmt.Println(arr)
}
