package main

import "fmt"

func main(){
	fmt.Println(CoinChange2(5,[]int{1,2,5}))
}

func CoinChange2(amount int, coins []int) int {
	combi := make([]int,amount+1)
	combi[0] = 1;
	for i := 0; i < len(coins); i++ {
		for j := 1; j < amount+1; j++{
			if j - coins[i] >= 0 {
				combi[j] = combi[j]+combi[j-coins[i]]
			}
		}
	}
	return combi[amount]
}