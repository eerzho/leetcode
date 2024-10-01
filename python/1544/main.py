class Solution:
    def makeGood(self, s: str) -> str:
        stack = []

        for c in s:
            if not stack:
                stack.append(c)
            elif c.lower() == stack[-1].lower():
                if c.islower() and not stack[-1].islower():
                    stack.pop()
                elif not c.islower() and stack[-1].islower():
                    stack.pop()
                else:
                    stack.append(c)
            else:
                stack.append(c)
            
        
        return ''.join(stack)