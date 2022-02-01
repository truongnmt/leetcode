func lengthOfLongestSubstring(s string) int {
	// Input: s = "abcabcbb" 3
	//              l   r
	//             a:0 b:1 c:2
	//     Output: 3
	//     "abc"

	//     Input: s = "pwwkew"
	//                    l r
	//                 ["w"]
	//                 ["pw"]
	//                 ["e"]
	//                 ["kew"]

	//     Input: s = "dvdf"
	//                 l r
	//         arr[i]  ["f", "vdf", "df"]

	//     Input: s = "pwwkew"
	//                 l r

	mapChar := make(map[rune]int)
	left, right := 0, 0
	max := 0
	for i, char := range s {
		// fmt.Printf("left=%v right=%v max=%v mapChar=%v\n", left, right, max, mapChar)
		if duplicatedIndex, isDuplicated := mapChar[char]; isDuplicated {
			if (right - left) > max {
				max = right - left
			}

			for j := left; j < duplicatedIndex; j++ {
				delete(mapChar, rune(s[j]))
			}

			left = duplicatedIndex + 1
			mapChar[char] = i
			right++
		} else {
			mapChar[char] = i
			right++
		}
	}

	if (right - left) > max {
		max = right - left
	}

	return max
}
