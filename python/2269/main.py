class Solution:
    def divisorSubstrings(self, num: int, k: int) -> int:
        digits = [int(digit) for digit in str(num)]

        begin = 0
        window_state = []

        result = 0
        for end in range(len(digits)):
            window_state.append(digits[end])
            if len(window_state) == k:
                window_condition = int(''.join(map(str, window_state)))
                if window_condition != 0 and num % window_condition == 0:
                    result += 1
                window_state.pop(0)
                begin += 1
        
        return result

solution = Solution()

# case 1
result = solution.divisorSubstrings(240, 2)
print(result)