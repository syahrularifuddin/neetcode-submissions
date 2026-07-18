class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        if len(nums) == 0 or len(nums) == 1:
            return
        i,j=0,1
        while j < len(nums):
            if nums[i] != 0:
                i=i+1
            if nums[i] == 0 and nums[j] != 0:
                temp = nums[i]
                nums[i] = nums[j]
                nums[j] = temp
                i=i+1
            j=j+1