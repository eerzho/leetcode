class Solution:
    def minLength(self, s: str) -> int:
        stack = []

        for c in s:
            if not stack:
                stack.append(c)
            elif stack[-1] == 'A' and c == 'B':
                stack.pop()
            elif stack[-1] == 'C' and c == 'D':
                stack.pop()
            else:
                stack.append(c)
