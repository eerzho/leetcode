from typing import List

class Solution:
    def longestOnes(self, nums: List[int], k: int) -> int:
        begin = 0
        window_state = 0
        result = float('-inf')

        for end in range(len(nums)):
            if nums[end] == 0:
                window_state += 1

            while window_state > k:
                result = max(result, end - begin)
                if nums[begin] == 0:
                    window_state -= 1
                begin += 1

        result = max(result, end - begin + 1)
        
        return result
    
solution = Solution()

# case 1
# result = solution.longestOnes([1,1,1,0,0,0,1,1,1,1,0], 2)
# print(result)

# case 2
result = solution.longestOnes([0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], 3)
print(result)

# case 3
result = solution.longestOnes([0,0,0,1], 4)
print(result)

# case 4
result = solution.longestOnes([1,1,0,1], 1)
print(result)