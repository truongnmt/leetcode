# nums1  1 2 3 4] 5 6 7
#              j
# nums1Right = mid - nums2Right -1 = 5 - 1 - 1 = 3

# nums2  1 2] 3
# mid    l i r
# nums2Right = m = 1

# merge  1 1 2 2 3 3 4 5 6 7
# median         3 3
# midpos         4 5

# total: 10
# mid: 5

# while true
# l=0;r=len(nums2)-1
# mid=(l+r)//2

# if nums2[m] > nums1[m+1]:
#     r = m-1
# elif nums1[m] > nums2[m+1]:

class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        A = nums1 if len(nums1) < len(nums2) else nums2
        B = nums2 if len(nums2) < len(nums1) else nums1
        
        total = len(nums1)+len(nums2)
        half = total // 2

        l, r = 0, len(A)-1
        while True:
            i = (l+r) // 2 # A
            j = half - i - 2 # B

            Aleft = A[i] if i >= 0 else float("-infinity")
            Bleft = B[j] if j >= 0 else float("-infinity")
            Aright = A[i+1] if (i+1) < len(A) else float("infinity")
            Bright = B[j+1] if (j+1) < len(B) else float("infinity")

            if Aleft <= Bright and Bleft <= Aright:
                if total%2 != 0:
                    return min(Aright, Bright)
                else:
                    return (max(Aleft, Bleft) + min(Aright, Bright)) / 2
            elif Aleft > Bright:
                r = i - 1
            else:
                l = i + 1
