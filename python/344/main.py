class Solution(object):
    def reverseString(self, s):
        left = 0
        right = len(s) - 1
        while left < right:
            s[left], s[right] = s[right], s[left]
            left += 1
            right -= 1

solution = Solution()

# case 1
s = list("hello")
solution.reverseString(s)
print(s)