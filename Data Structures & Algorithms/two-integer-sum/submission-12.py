class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
      cmap = {}
      res = []
      for i,n in enumerate(nums):
        cmap[n] = i
      for i,n in enumerate(nums):
        if target-n in cmap:
          v1=i
          idx = cmap.get(target-n)
          if i == idx:
              continue
          v2=idx
          return [v1,v2]
      