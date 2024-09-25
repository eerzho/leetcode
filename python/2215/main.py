from typing import List
from collections import defaultdict


class Solution:
    def findDifference(self, nums1: List[int], nums2: List[int]) -> List[List[int]]:
        hash_map1 = defaultdict(int)     
        hash_map2 = defaultdict(int)

        for num in nums1:
            hash_map1[num] += 1

        for num in nums2:
            hash_map2[num] += 1

        result1 = []
        for num in nums1:
            if hash_map2[num] == 0:
                if hash_map1[num] == 1:
                    result1.append(num)
                else:
                    hash_map1[num] -= 1

        result2 = []
        for num in nums2:
            if hash_map1[num] == 0:
                if hash_map2[num] == 1:
                    result2.append(num)
                else:
                    hash_map2[num] -= 1

        return [result1, result2]

            
solution = Solution()

result = solution.findDifference([1,2,3,3], [1,1,2,2])
print(result)