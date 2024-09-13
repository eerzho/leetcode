from typing import List

class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        left = 0
        right = len(numbers) - 1
        while left < right:
            sum = numbers[left] + numbers[right]
            if sum > target:
                right -= 1
            elif sum < target:
                left += 1
            elif sum == target:
                return [left + 1, right + 1]
        return []
    
solution = Solution()

# case 1
numbers = [2,7,11,15]
target = 9
result = solution.twoSum(numbers, target)
print(result)

# case 2
numbers = [2,3,4]
target = 6
result = solution.twoSum(numbers, target)
print(result)

# case 1
numbers = [-1,0]
target = -1
result = solution.twoSum(numbers, target)
print(result)