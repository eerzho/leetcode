from typing import List, Deque

class Solution:
    def timeRequiredToBuy(self, tickets: List[int], k: int) -> int:
        i = 0
        result = 0
        while True:
            if tickets[i] != 0:
                tickets[i] -= 1
            
            if tickets[k] == 0:
                return result
            
            i += 1
            result += 1
            if i == len(tickets):
                i = 0