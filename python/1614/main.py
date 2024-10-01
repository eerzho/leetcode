class Solution:
    def maxDepth(self, s: str) -> int:
        stack = []

        result = 0
        for c in s:
            if not stack and (c == '(' or c == ')'):
                stack.append(c)
            else:
                if c == "(":
                    stack.append(c)
                elif c == ")" and stack[-1] == "(":
                    stack.pop()
            result = max(result, len(stack))

        return result

