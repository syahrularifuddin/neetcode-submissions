class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        count = 0
        s=s.rstrip()
        for i in reversed(s):
            if i == " ":
                break
            count += 1
        return count