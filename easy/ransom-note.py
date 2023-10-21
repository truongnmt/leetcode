class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        # O(m+n)
        ransomNoteMap = defaultdict(int)
        magazineMap = defaultdict(int)

        for ch in ransomNote:
            ransomNoteMap[ch] += 1
        for ch in magazine:
            magazineMap[ch] += 1
        
        for k, v in ransomNoteMap.items():
            if k not in magazineMap or magazineMap[k] < v:
                return False
        
        return True

        # O(m*n)
        # surprisingly in python, this run faster?
        for letter in ransomNote:
            if magazine.find(letter) != -1:
                magazine = magazine.replace(letter, '', 1)
            else:
                return False

        return True
