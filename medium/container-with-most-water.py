class Solution:
    def maxArea(self, height: List[int]) -> int:
        # brute force
        # maxWater = 0
        
        # for l in range(len(height)):
        #     for r in range(l+1, len(height)):
        #         area = min(height[l], height[r])*(r-l)
        #         maxWater = max(maxWater, area)

        # return maxWater

        # O(n)
        # two pointer
        maxWater = 0
        l, r = 0, len(height)-1

        while l<r:
            area = min(height[l], height[r])*(r-l)
            maxWater = max(maxWater, area)

            if height[l] < height[r]:
                l += 1
            else:
                r -= 1
        
        return maxWater
