class Solution:
    def sumOfUnique(self, nums: List[int]) -> int:
        # numsMap = {}
        numsMap = [0]*101 
        # Generally array is faster than hash map 
        # because it doesn't have overhead like hashing, resolving collision
        sum = 0

        for n in nums:
            # print(f"n={n} numsMap[n]={numsMap[n]}")
            if numsMap[n] > 0:
                numsMap[n] += 1
                if numsMap[n] == 2:
                    sum -= n
            else:
                sum += n
                numsMap[n] = 1

        return sum
