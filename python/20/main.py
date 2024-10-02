class Solution:
    def isValid(self, s: str) -> bool:
        pair = {"(": ")", "{": "}", "[": "]"}

        stack = []
        for c in s:
            if c in pair:
                stack.append(c)
            else:
                if len(stack) == 0:
                    return False

                prev = pair[stack.pop()]
                if prev != c:
                    return False

        return len(stack) == 0
