class Solution:
    def findMissingRanges(self, nums: List[int], lower: int, upper: int) -> List[List[int]]:
      result = []
      if not nums:
        return [[lower, upper]]
      for i in range(len(nums)):
        if i == 0: # first
          if nums[i] - lower >= 1:
            result.append([lower, nums[i]-1])
        else:
          if nums[i] - nums[i-1] >= 2:
            print(nums[i], nums[i-1])
            result.append([nums[i-1]+1, nums[i]-1])
      if upper-nums[-1] >=1:
        result.append([nums[-1]+1,upper])
      return result