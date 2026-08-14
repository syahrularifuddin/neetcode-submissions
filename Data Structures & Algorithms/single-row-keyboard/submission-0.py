class Solution:
    def calculateTime(self, keyboard: str, word: str) -> int:
        index = {}
        for k,v in enumerate(keyboard):
          index[v]=k
        dist = 0
        last = 0
        for w in word:
          dist += abs(last-index[w])
          last = index[w]
        return dist