from typing import List
from collections import defaultdict

class Solution:
    def uniqueOccurrences(self, arr: List[int]) -> bool:
        hashMap1 = defaultdict(int)

        for num in arr:
            hashMap1[num] += 1
        
        hashMap2 = defaultdict(int)
        for num in hashMap1:
            hashMap2[hashMap1[num]] += 1
        
        if len(hashMap1) == len(hashMap2):
            return True
        else:
            return False

solution = Solution()

result = solution.uniqueOccurrences([-3,0,1,-3,1,1,1,-3,10,0])
