class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
      if len(s) != len(t):
        return False
      sorted_a = ''.join(sorted(s))
      sorted_b = ''.join(sorted(t))
      return sorted_a == sorted_b