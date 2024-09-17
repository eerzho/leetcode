from typing import List
import sys

class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        begin = 0
        sum = 0
        result = -sys.maxsize - 1

        for end in range(0, len(nums)):
            sum += nums[end]
            if end >= k - 1:
                result = max(result, sum / k)
                sum -= nums[begin]
                begin+=1

        return result
    
solution = Solution()

# case 1
result = solution.findMaxAverage([-1], 1)
print(result)

# case 2
result = solution.findMaxAverage([1,12,-5,-6,50,3], 4)
print(result)

# case 3
result = solution.findMaxAverage([0,4,0,3,2], 1)
print(result)