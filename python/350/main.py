from typing import List
from collections import defaultdict

class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        hash_map = defaultdict(int)
        for num in nums1:
            hash_map[num] += 1

        result = []
        for num in nums2:
            if hash_map[num] > 0:
                result.append(num)
                hash_map[num] -= 1
        
        return result
    
solution = Solution()

result = solution.intersect([1,2,2,1], [2,2])
print(result)