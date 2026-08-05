class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        count = 0
        s=s.rstrip()
        print(s)
        for i in reversed(s):
            if i == " ":
                break
            print(i)
            count += 1
        return count