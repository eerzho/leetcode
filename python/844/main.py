class Solution:
    def backspaceCompare(self, str1: str, str2: str) -> bool:
        return self.cleanStr(str1) == self.cleanStr(str2)
    def cleanStr(self, s: str) -> str:
        str_r = ""
        str_ind = len(s) - 1
        counter = 0
        while str_ind >= 0:
            if s[str_ind] == "#":
                counter += 1
            elif counter > 0 and s[str_ind] != "#":
                counter -= 1
            elif counter == 0:
                str_r += s[str_ind]
            str_ind -= 1
        return str_r


solution = Solution()

# case 1
result = solution.backspaceCompare('"ab#c"', '"ad#c"')
print(result)