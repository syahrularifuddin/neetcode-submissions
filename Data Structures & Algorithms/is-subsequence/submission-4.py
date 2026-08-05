class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
      c = 0
      print(len(s))
      if len(s)==0:
        return True
      for i in t:
        if c == len(s):
          return True
        if i == s[c]:
          print(i,s[c])
          c=c+1
      print(c)
      return c == len(s)