# 0 1 3 5 5
#     i
# 0 0 2 4 4
# ops = 1

class Solution:
    def minimumOperations(self, nums: List[int]) -> int:
        # attemp1: brute force O(n^2)
        # nums.sort()
        # if nums[len(nums)-1]==0:
        #     return 0
        # ops = 0
        # for i in range(len(nums)):
        #     if nums[i]==0:
        #         continue
        #     ops += 1
        #     diff = nums[i]
        #     for j in range(i, len(nums)):
        #         nums[j] = max(0, nums[j]-diff)
        # return ops



        # attemp2: count unique item O(n)
        return len(set(nums)-{0})
