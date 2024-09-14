from typing import List

class Solution:
    def maxArea(self, height: List[int]) -> int:
        l = 0
        r = len(height)        
        result = 0
        while l < r:
            dif = r - l
            minH = min(height[l], height[r])
            result = max(result, dif * minH)
            if height[l] < height[r]:
                l += 1
            else:
                r -= 1
        return result