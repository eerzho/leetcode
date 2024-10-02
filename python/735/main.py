from typing import List


class Solution:
    def asteroidCollision(self, asteroids: List[int]) -> List[int]:
        stack = []

        for a in asteroids:
            if not stack:
                stack.append(a)
            elif stack[-1] < 0 and a < 0:
                stack.append(a)
            elif stack[-1] >= 0 and a >= 0:
                stack.append(a)
            else:
                while stack and stack[-1] > 0 and stack[-1] < abs(a):
                    stack.pop()

        return stack


solution = Solution()

result = solution.asteroidCollision([-2, -2, 1, -2])
print(result)
