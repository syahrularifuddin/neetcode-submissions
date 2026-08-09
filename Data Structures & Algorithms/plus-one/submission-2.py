class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        t = 0
        for d in digits:
            t = (t*10)+d
        t+=1
        res = []
        while t>0:
            res.insert(0, t%10)
            t=t//10
        return res