from typing import List

class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        begin = 0
        window_state = 0

        nums_sum = 0
        for i in range(len(nums)):
            nums_sum += nums[i]

        if nums_sum == target:
            return len(nums)
        elif nums_sum < target:
            return 0
        
        result = len(nums)
        for end in range(len(nums)): 
            window_state += nums[end]
            while window_state >= target:
                result = min(result, end - begin + 1)
                window_state -= nums[begin]
                begin += 1

        
        return result


solution = Solution()

# case 1
result = solution.minSubArrayLen(7, [2,3,1,2,4,3])
print(result)

# case 2
result = solution.minSubArrayLen(4, [1,4,4])
print(result)

# case 3
result = solution.minSubArrayLen(11, [1,2,3,4,5])
print(result)

# case 4
result = solution.minSubArrayLen(15, [1,2,3,4,5])
print(result)