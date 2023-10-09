# 4 5 6 7
#         0 1 2      

class Solution:
    def search(self, nums: List[int], target: int) -> int:
        # find rotated position
        l = 0; r = len(nums)-1
        while (r-l)>1:
            mid = (l+r)//2
            if nums[0] <= nums[mid]:
                l = mid
            else:
                r = mid

        pivot = l
        l = 0; r = len(nums)-1
        
        # finding which sorted array part that the target is in
        if nums[0] <= target <= nums[pivot]:
            r = pivot
        else:
            l = pivot+1

        # binary search
        while l<=r:
            mid = (l+r)//2
            if nums[mid] == target:
                return mid

            if nums[mid] < target:
                l = mid+1
            else:
                r = mid-1
        
        return -1
