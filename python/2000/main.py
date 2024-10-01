class Solution:
    def reversePrefix(self, word: str, ch: str) -> str:
        word_list = [x for x in word]

        for i in range(len(word_list)):
            if word_list[i] == ch:
                left = 0
                right = i
                while left < right:
                    word_list[left], word_list[right] = word_list[right], word_list[left]

        return ''.join(word_list)