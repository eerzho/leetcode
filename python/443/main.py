from typing import List

class Solution:
    def compress(self, chars: List[str]) -> int:
        if len(chars) < 2:
            return len(chars)

        begin = 0

        result = []
        count = 1
        for end in range(1, len(chars)):
            if chars[end] == chars[begin]:
                count += 1
            else:
                result.append(chars[begin])
                if count != 1:
                    list1 = [x for x in str(count)]
                    for i in range(len(list1)):
                        result.append(list1[i])
                count = 1
                begin = end

        if count > 0:
            result.append(chars[begin])
            if count != 1:
                list1 = [x for x in str(count)]
                for i in range(len(list1)):
                    result.append(list1[i])

        for i in range(len(result)):
            chars[i] = result[i]

        return len(result)