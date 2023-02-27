package others

import (
	"fmt"
	"math"
)

//Maximize the profit by selecting a single day to buy one stock and selecting a different day in the future to sell that stock.

func maxProfit(prices []int) int {
    lenPrices := len(prices)
    if lenPrices == 0 || lenPrices == 1 {
        return 0
    }
    lowestPriceIdx := 0
    maxProfit := 0
    for i := 1; i<lenPrices; i++ {
        if prices[i] < prices[lowestPriceIdx] {
            lowestPriceIdx = i
            continue
        }
        maxProfit = int(math.Max(float64(maxProfit), float64(prices[i] - prices[lowestPriceIdx])))
    }
    return maxProfit
}

func Run_Buy_and_Sell_Stock() {
	fmt.Println("MaxProfit for given Prices array::",[]int{7,1,5,3,6,4},"is", maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println("MaxProfit for given Prices array::",[]int{7,6,5,4,3,2},"is", maxProfit([]int{7,6,5,4,3,2}))
}