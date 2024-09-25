from typing import List

class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        hash_map = {}
        for num in nums:
            hash_map[num] = True

        i = 0
        while i < len(nums):
            if hash_map.get(i) == None:
                return i
            i += 1
        return i


solution = Solution()

result = solution.missingNumber([3,0,1])
print(result)