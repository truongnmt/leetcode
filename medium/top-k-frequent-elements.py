# build the character freq hash map
# sort the hash map by value in descending order
# append result considering input k

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        if k==len(nums):
            return nums

        freq = defaultdict(int)

        for num in nums:
            freq[num] += 1

        freq_item = list(freq.items()) # [(1, 3), (2, 2), (3, 1)]
        freq_item.sort(key=lambda item: item[1], reverse=True)
        print(freq_item)

        result = []
        for key, value in freq_item:
            if len(result)<k:
                result.append(key)

        return result
