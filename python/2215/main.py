from typing import List, Dict
from collections import defaultdict


class Solution:
    def findDifference(self, nums1: List[int], nums2: List[int]) -> List[List[int]]:
        def find(hash_map: Dict, nums: List[int]) -> List[int]:
            result = []
            for num in nums:
                if hash_map.get(num) == None:
                    hash_map[num] = False
                    result.append(num)
            return result
        
        def create(nums: List[int]) -> Dict:
            hash_map = {}
            for num in nums:
                hash_map[num] = True
            return hash_map

        return [find(create(nums2), nums1), find(create(nums1), nums2)]


            
solution = Solution()

result = solution.findDifference([1,2,3,3], [1,1,2,2])
print(result)