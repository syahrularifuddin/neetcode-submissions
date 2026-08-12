class Solution:
    def confusingNumber(self, n: int) -> bool:
        pair = {
          0: 0,
          1: 1,
          6: 9,
          8: 8,
          9: 6,
        }
        init = n
        nums = []
        while n > 0:
          nums.append(n%10)
          n//=10
        print(1, nums)
        new_nums = []
        for num in nums:
          if num not in pair:
            return False
          new_nums.append(pair.get(num))
        print(2, new_nums)
        new_n = 0
        for new_num in new_nums:
          new_n = new_n * 10
          new_n += new_num
        print(3, new_n, init, new_n != init)
        return init != new_n
          