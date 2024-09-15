class Solution:
    def reverseVowels(self, s: str) -> str:
        ss = [x for x in s]

        left = 0
        right = len(s) - 1

        while left < right: 
            if not self.isVowel(ss[left]): 
                left += 1
            elif not self.isVowel(ss[right]):
                right -= 1
            else:
                ss[left], ss[right] = ss[right], ss[left]
                left += 1
                right -= 1

        return "".join(ss)
    
    def isVowel(self, s: str) -> bool:
        s = s.lower()
        if s == 'a':
            return True
        elif s == 'e':
            return True
        elif s == 'i':
            return True
        elif s == 'o':
            return True
        elif s == 'u':
            return True
        else:
            return False
        
solution = Solution()

# case 1
result = solution.reverseVowels("IceCreAm")
print(result)