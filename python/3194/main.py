from typing import List

class Solution:
    def minimumAverage(self, nums: List[int]) -> float:
        nums.sort()

        left = 0
        right = len(nums) - 1
        result = nums[right]

        while left < right:
            result = min(result, (nums[left] + nums[right]) / 2)
            left += 1
            right -= 1
        
        return result