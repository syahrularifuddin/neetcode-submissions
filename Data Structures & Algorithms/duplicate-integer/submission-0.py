class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        cmap = {}
        for n in nums:
          cmap[n] = cmap.get(n, 0)+1
        print(cmap)
        for c in cmap.values():
          if c > 1:
            return True
        return False