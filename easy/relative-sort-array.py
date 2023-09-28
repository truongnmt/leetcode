class Solution:
    def relativeSortArray(self, arr1: List[int], arr2: List[int]) -> List[int]:

        # items = [(2, 2), (3, 1), (1, 3)]
        # item.sort(key=lambda item: item[1])
        # sort the items with the key is the item value, item[1]
        # items = [(3, 1), (2, 2), (1, 3)]

        pos = {x: i for i, x in enumerate(arr2)}
        print(pos)
        # {2: 0, 1: 1, 4: 2, 3: 3, 9: 4, 6: 5}
        # when sorting each value, find its position by lookup in arr2
        # the key is the value index
        # if the key not found, increment the number with the len(arr2)

        return sorted(arr1, key=lambda x: pos.get(x, x+len(arr2)))
        