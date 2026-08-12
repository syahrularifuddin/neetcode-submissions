class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        profit = 0
        buy = prices[0]
        for price in prices:
          diff = price - buy
          profit = max(profit, diff)
          if price < buy:
            buy = price
        return profit