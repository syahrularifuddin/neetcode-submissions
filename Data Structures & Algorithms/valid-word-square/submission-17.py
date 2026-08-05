class Solution:
    def validWordSquare(self, words: List[str]) -> bool:
      for i in range(len(words)):
        col = ""
        for j in range(len(words)):
          if j >= len(words[i]) or i >= len(words[j]):
            print("broken")
            break
          print(j,i)
          col += words[j][i]
        print(col)
        if words[i] != col:
          return False
      return True