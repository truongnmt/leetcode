# build the character freq hash map
# sort the hash map by value in descending order
# append result considering input k

# nums = [1,1,1,2,2,3], k = 2
#         i
#       [(1, 3), (2, 2), (3, 1)]
# freq    1 2 3 4 5 6
#         3 2 1  

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        if k==len(nums):
            return nums

        freq = defaultdict(int)
        for num in nums:
            freq[num] += 1

        # attemp1: sort the hashmap
        # freq_item = list(freq.items()) # [(1, 3), (2, 2), (3, 1)]
        # freq_item.sort(key=lambda item: item[1], reverse=True)

        # result = []
        # for key, value in freq_item:
        #     if len(result)<k:
        #         result.append(key)

        # attemp2: build list of frequence, with index is the frequence
        sort_freq = [[] for i in range(len(nums)+1)] 
        # create list of empty lists [[], [], [], [], ...]
        # while result = []*26 create empty array with length = 26

        for key, value in list(freq.items()):
            sort_freq[value].append(key)
        
        result = []
        for i in range(len(sort_freq)-1, -1, -1):
            if len(result)<k:
                result += sort_freq[i]

        return result
