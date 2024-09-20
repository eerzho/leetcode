class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        list1 = [x for x in word1]
        list2 = [x for x in word2]

        p1 = 0
        p2 = 0

        result = []
        while p1 < len(list1) and p2 < len(list2):
            if (p1 + p2) % 2 == 0:
                result.append(list1[p1])
                p1 += 1
            else:
                result.append(list2[p2])
                p2 += 1
        
        while p1 < len(list1):
            result.append(list1[p1])
            p1 += 1

        while p2 < len(list2):
            result.append(list2[p2])
            p2 += 1
            
        return "".join(result)
    

