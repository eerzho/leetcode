from typing import List

class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        p1 = 0
        p2 = 0

        result = []

        while p1 < m and p2 < n:
            if nums1[p1] > nums2[p2]:
                result.append(nums2[p2])
                p2 += 1
            else:
                result.append(nums1[p1])
                p1 += 1
        
        for i in range(p1, m):
            result.append(nums1[i])

        for i in range(p2, n):
            result.append(nums2[i])

        for i in range(len(nums1)):
            nums1[i] = result[i]


solution = Solution()

# case 1
nums1 = [1,2,3,0,0,0] 
m = 3 
nums2 = [2,5,6] 
n = 3
# solution.merge(nums1, m, nums2, n)
# print(nums1)


# case 2
nums1 = [4,0,0,0,0,0]
m = 1
nums2 = [1,2,3,5,6]
n = 5
# solution.merge(nums1, m, nums2, n)
# print(nums1)

# case 3
nums1 = [1]
m = 1
nums2 = []
n = 0
solution.merge(nums1, m, nums2, n)
print(nums1)