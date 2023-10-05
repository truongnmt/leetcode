# 8 9 7 9
# 7 8 9 9

# 3 9 1 1 5 8
# 1 1 3 5 8 9

class Solution:
    def maxWidthOfVerticalArea(self, points: List[List[int]]) -> int:
        # x_cordinates = [x for x, _ in points]
        # x_cordinates.sort()
        arr = sorted(x for x,_ in points)
        
        diff = 0
        for i in range(1, len(x_cordinates)):
            diff = max(x_cordinates[i]-x_cordinates[i-1], diff)

        return diff
