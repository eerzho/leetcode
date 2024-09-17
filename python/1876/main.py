from collections import defaultdict

class Solution:
    def countGoodSubstrings(self, s: str) -> int:
        s_list = [x for x in s]

        begin = 0
        window_state = defaultdict(int)

        result = 0
        for end in range(len(s_list)):
            window_state[s_list[end]] += 1

            if end - begin + 1 == 3:
                skip = False
                for i in range(begin, end + 1):
                    if window_state[s_list[i]] > 1:
                        skip = True
                        continue
                if not skip:
                    result += 1
                window_state[s_list[begin]] -= 1
                begin += 1
                
            
        return result
    
solution = Solution()

# case 1
result = solution.countGoodSubstrings("xyzzaz")
print(result)

# case 2
result = solution.countGoodSubstrings("owuxoelszb")
print(result)

# case 3
result = solution.countGoodSubstrings("aababcabc")
print(result)