# nums1 = [1,2,3,0,0,0], m = 3, 
#          m     
#                  i
#          1 2 2 3 5 6
# nums2 = [2,5,6], n = 3
#        n

class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """

        # O( (m+n)log(m+n) )
        # nums1[m:] = nums2
        # nums1.sort()

        # O(m+n)
        while m>0 and n>0:
            if nums1[m-1] >= nums2[n-1]:
                nums1[m+n-1] = nums1[m-1]
                m -= 1
            else:
                nums1[m+n-1] = nums2[n-1]
                n -= 1
        
        if n>0:
            nums1[:n] = nums2[:n]
