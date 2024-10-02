class Solution:
    def removeDuplicates(self, s: str) -> str:
        stack = []

        for c in s:
            if not stack:
                stack.append(c)
            else:
                prev = stack.pop()
                if prev != c:
                    stack.append(prev)
                    stack.append(c)

        return "".join(stack)
