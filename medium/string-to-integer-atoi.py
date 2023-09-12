class Solution:
    def myAtoi(self, s: str) -> int:

        # read input
        # remove white space
        # read the sign -> positive / negative
        #               -> if next char is not digit, return 0
        # read digits -> convert to number
        # check negative
        # check in range [-2^31, 2^31 - 1]

        is_negative = False
        num = 0
        i = 0
        n = len(s)

        while (i<n and s[i] == " "):
            i += 1

        if (i<n and s[i] == "-"):
            is_negative = True
            if (i+1<n and not s[i+1].isdigit()):
                return 0
            i += 1
        if (i<n and s[i] == "+"):
            if (i+1<n and not s[i+1].isdigit()):
                return 0
            i += 1
        
        while (i<n and s[i].isdigit()):
            num = num*10 + ord(s[i]) - ord('0')
            if num > 2**31:
                break
            i += 1

        if (is_negative): num = num*-1

        num = min(num, 2**31-1)
        num = max(num, -2**31)

        return num
