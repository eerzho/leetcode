from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hash = {}

        for i in range(len(nums)):
            if nums[i] in hash:
                return [i, hash[nums[i]]]

            hash[target-nums[i]] = i

        return []
    

solution = Solution()

# case 1
result = solution.twoSum([2,7,11,15], 9)
print(result)