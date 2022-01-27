func groupAnagrams(strs []string) [][]string {
	//     Input: strs = ["eat","tea","tan","ate","nat","bat"]
	//     Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

	//     loop str in strs
	//     loop char in str
	//     get int char and put in array k
	//     put array k in map[[26]int][]string

	//     mapStr := make(map[[26]int][]string)
	//     for _, str := range strs {
	//         k := [26]int{}
	//         for _, character := range str {
	//             k[int(character)-int('a')] += 1
	//         }
	//         mapStr[k] = append(mapStr[k], str)
	//     }

	//     ans := [][]string{}
	//     for _, str := range mapStr {
	//         ans = append(ans, str)
	//     }
	//     return ans

	// using sort
	mapStr := make(map[string][]string)
	for _, str := range strs {
		cloneStr := str
		cloneStr = SortCharInString(cloneStr)
		mapStr[cloneStr] = append(mapStr[cloneStr], str)
	}

	ans := [][]string{}
	for _, str := range mapStr {
		ans = append(ans, str)
	}
	return ans
}

func SortCharInString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
