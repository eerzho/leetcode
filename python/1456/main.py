class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        
        def isVowel(s: str):
            if s == 'a' or s == 'e' or s == 'i' or s == 'o' or s == 'u':
                return True
            else:
                return False
        
        s_list = [x for x in s]
        begin = 0
        window_state = 0

        result = 0
        for end in range(len(s_list)):
            if isVowel(s_list[end]):
                window_state += 1

            if end - begin == k:
                if isVowel(s_list[begin]):
                    window_state -= 1
                begin += 1
            
            result = max(result, window_state)

        return result
    
solution = Solution()

# case 1
result = solution.maxVowels("abciiidef", 3)
print(result)
