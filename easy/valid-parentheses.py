class Solution:
    def isValid(self, s: str) -> bool:
        if len(s)%2 != 0:
            return False

        d = { '(':')', '{':'}', '[':']'}
        stack = []
        for i in s:
            # if it's the open bracket, append to stack
            if i in d:
                stack.append(i)
            # else, it's close bracket, check stack size and check matching
            else:
                if len(stack)==0 or d[stack.pop()] != i:
                    return False
        
        # check the stack still contains unclosed bracket
        return len(stack)==0
