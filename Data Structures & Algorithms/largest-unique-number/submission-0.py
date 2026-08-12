class Solution:
    def largestUniqueNumber(self, nums: List[int]) -> int:
        count = {}
        for n in nums:
          count[n] = count.get(n, 0)+1
        print(count)
        max_num = -1
        for k,v in count.items():
          if v > 1:
            continue
          if k > max_num:
            max_num = k
        return max_num