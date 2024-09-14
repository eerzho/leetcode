class Solution:
    def isSubsequence(self, sub: str, s: str) -> bool:
        sub_ind = 0
        for i in range(len(s)):
            if sub_ind == len(sub):
                return True
            if s[i] == sub[sub_ind]:
                sub_ind += 1
            
        return sub_ind == len(sub)