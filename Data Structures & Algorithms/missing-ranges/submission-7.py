class Solution:
    def findMissingRanges(self, nums: List[int], lower: int, upper: int) -> List[List[int]]:
      result = []
      if len(nums) == 0:
        result.append([lower,upper])
        return result
      for i in range(len(nums)):
        if i == 0: # first
          if nums[i] - lower >= 1:
            result.append([lower, nums[i]-1])
        else:
          if nums[i] - nums[i-1] >= 2:
            print(nums[i], nums[i-1])
            result.append([nums[i-1]+1, nums[i]-1])
      if upper-nums[len(nums)-1] >=1:
        result.append([nums[len(nums)-1]+1,upper])
      return result