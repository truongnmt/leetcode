class Solution {
    fun maxProfit(prices: IntArray): Int {
        var maxProfit = 0
        var minPriceIndex = 0

        for (i in prices.indices) {
            if (prices[minPriceIndex] > prices[i]) {
                minPriceIndex = i
            }

            if ( (prices[i]-prices[minPriceIndex])>maxProfit ) {
                maxProfit = prices[i]-prices[minPriceIndex]
            }
        }

        return maxProfit
    }
}
