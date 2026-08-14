class Solution:
    def canJump(self, nums: List[int]) -> bool:
      furthest = 0
      for i in range(len(nums)):
          if i > furthest:
              return False  # Can't even reach this index
          furthest = max(furthest, i + nums[i])
      return True