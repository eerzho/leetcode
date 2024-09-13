from typing import List

class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        n = len(nums)
        left = 0
        right = n - 1
        result = [0] * n

        for i in range(n - 1, -1, -1):
            leftN = nums[left]**2
            rightN = nums[right]**2
            if leftN > rightN:
                result[i] = leftN
                left += 1
            else:
                result[i] = rightN
                right -= 1
        return result

solution = Solution()

nums = [-4,-1,0,3,10]
res = solution.sortedSquares(nums)
print(res)