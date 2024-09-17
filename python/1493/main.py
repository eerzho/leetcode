from typing import List

class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        begin = 0
        window_state = 0
        k = 1
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
        
        return result - 1
    

solution = Solution()

# case 1
result = solution.longestSubarray([1,1,0,1])
print(result)