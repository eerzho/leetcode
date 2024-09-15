class Solution:
    def validPalindrome(self, s: str) -> bool:
        ss = [x for x in s]

        left = 0
        right = len(ss) - 1
        status = False

        l_status = True
        while left < right:
            if ss[left] != ss[right]:
                if status == True:
                    l_status = False
                    break
                else:
                    left += 1
                    status = True
                    continue
            left += 1
            right -= 1
        
        left = 0
        right = len(ss) - 1
        status = False

        r_status = True
        while left < right:
            if ss[left] != ss[right]:
                if status == True:
                    r_status = False
                    break
                else:
                    right -= 1
                    status = True
                    continue
            left += 1
            right -= 1
            
        return l_status or r_status
                

solution = Solution()

# case 1
result = solution.validPalindrome("axbcbaba")
print(result)

# case 2 
result = solution.validPalindrome("cbbcc")
print(result)

# case 3
result = solution.validPalindrome("abc")
print(result)