class Solution(object):
    def isPalindrome(self, s):
        """
        :type s: str
        :rtype: bool
        """
        left = 0
        right = len(s) - 1
        while left < right:
            if not s[left].isalnum():
                left += 1
                continue
            if not s[right].isalnum():
                right -= 1
                continue
            if s[left].lower() == s[right].lower():
                left += 1
                right -= 1
            else:
                return False
        return True
    

solution = Solution()

# case 1
s = "A man, a plan, a canal: Panama"
result = solution.isPalindrome(s)
print(result)

# case 2
s = "race a car"
result = solution.isPalindrome(s)
print(result)

# case 3
s = " "
result = solution.isPalindrome(s)
print(result)

