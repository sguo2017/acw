/**
使我们感兴趣的点是上图中的峰和谷。我们需要找到最小的谷之后的最大的峰。
我们可以维持两个变量——minprice 和 maxprofit，它们分别对应迄今为止所得到的最小的谷值和最大的利润（卖出价格与最低价格之间的最大差值）。

复杂度分析
时间复杂度：O(n)。只需要遍历一次。
空间复杂度：O(1)。只使用了两个变量。
*/
package main

import "fmt"

const INT_MAX = int(^uint(0) >> 1)

func main() {
	fmt.Printf("max profit\n")
	var cs = []int{7, 1, 5, 3, 6, 4}
	fmt.Print(maxProfit(cs))
}

func maxProfit(candlestick []int) int {
	minprice := INT_MAX
	maxProfit := 0
	var i int
	l := len(candlestick)
	for i = 1; i < l; i++ {
		if candlestick[i] < minprice {
			minprice = candlestick[i]
		}

		if candlestick[i]-minprice > maxProfit {
			maxProfit = candlestick[i] - minprice
		}
	}
	return maxProfit
}
