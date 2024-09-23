from typing import List

class Solution:
    def maxOperations(self, nums: List[int], k: int) -> int:
        nums.sort()
        left = 0
        right = len(nums) - 1

        result = 0
        while left < right: 
            sum = nums[left] + nums[right]
            if sum == k:
                left += 1 
                right -= 1
                result += 1
            elif sum > k:
                right -= 1
            else:
                left += 1

        return result


solution = Solution()

# case 1 
result = solution.maxOperations([3,1,3,4,3], 6)
print(result)