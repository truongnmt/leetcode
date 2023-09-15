# [100,4,200,1,3,2]
#  i

# set()
# min, max
# for num in range(min..max):
#     if num and num+1 exists in set
#         count++


class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        if len(nums) == 0: return 0

        nums_set = set()
        min_n = max_n = nums[0]

        for num in nums:
            nums_set.add(num)
            min_n = min(min_n, num)
            max_n = max(max_n, num)

        max_count = count = 1
        for num in nums_set:
            
            if num-1 not in nums_set:
                count = 1

                while num+count in nums_set:
                    count += 1
                max_count = max(max_count, count)

        return max_count
