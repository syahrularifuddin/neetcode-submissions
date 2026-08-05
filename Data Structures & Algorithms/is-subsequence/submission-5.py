class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
      c = 0
      if len(s)==0:
        return True
      for i in t:
        if c == len(s):
          return True
        if i == s[c]:
          c=c+1
      return c == len(s)