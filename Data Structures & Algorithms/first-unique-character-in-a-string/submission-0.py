class Solution:
    def firstUniqChar(self, s: str) -> int:
        count = {}
        for c in s:
            count[c] = count.get(c, 0) + 1
        print(count)
        for k,v in count.items():
            if v == 1:
                return s.find(k)
        return -1
