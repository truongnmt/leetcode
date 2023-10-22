# abacaba

# ab | a
# ac | a
# loop through s, use set to check ch exist in sub

class Solution:
    def partitionString(self, s: str) -> int:
        # count = [0]*26
        # total = 1

        # for ch in s:
        #     index = ord(ch)-ord('a')
        #     count[index] += 1
        #     if count[index] > 1:
        #         total += 1
        #         count = [0]*26
        #         count[index] = 1

        # return total

        sub = ""
        l = 0
        for ch in s:
            if ch not in sub:
                sub += ch
            else:
                l += 1
                sub = ch
        return l + 1
