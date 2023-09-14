#"eat","tea","tan","ate","nat","bat"
#
# create a hash map with a key is lists of character count, values are string
# [1e 1a 1t] eat aet tea
# use an int array with the size of 26 as key

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        res = defaultdict(list) 
        # # mapping charCount to list of Anagrams
        # # defaultdict is a dict, hash map with a default value of an empty list

        for str in strs:
            count = [0]*26
            for ch in str:
                count[ord(ch)-ord("a")] += 1
            res[tuple(count)].append(str)
            # use tuple because tuple is hashable, immutable, it can be used as a key to a hash map
            # list can't be a key in hash map
        return res.values()

        # result = {}
        # for str in strs:
        #     count = [0]*26
        #     for ch in str:
        #         count[ord(ch)-ord("a")] += 1    
        #     if tuple(count) in result:
        #         result[tuple(count)].append(str)
        #     else:
        #         result[tuple(count)] = [str]
        # return result.values()
