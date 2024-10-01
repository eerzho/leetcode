from typing import List


class Solution:
    def minOperations(self, nums: List[int]) -> int:
        p1 = 0
        p2 = p1 + 2

        while p2 < len(nums):
            for i in range(p1, p2 + 1):
                print(i)
            p1 += 1
            p2 += 1
