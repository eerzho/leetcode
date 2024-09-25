from typing import List

class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        hash_map = {}        

        for num in nums1:
            hash_map[num] = True

        result = []
        for num in nums2:
            if hash_map.get(num) == True:
                result.append(num)
                hash_map[num] = False
                

        return result