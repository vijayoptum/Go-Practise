// when we need to find the number of coins to get the amount and can take as coins for each denomination

package main

func coingchange(coins []int, length int, amount int) int {

	if amount == 0 {
		return 1
	}
	if amount < 0 {
		return 0
	}
	if length <= 0 && amount >= 1 {
		return 0
	}

	return coingchange(coins, length-1, amount) + coingchange(coins, length, amount-coins[length-1])
}

func main() {
	amount := 10
	coins := []int{2, 5, 3, 6}

	numberOfCoins := coingchange(coins, len(coins), amount)
	println(numberOfCoins)

}
