// s1 = "ab", s2 = "eidbaooo"
//  ab
// eidbaooo

// build map of occurence in s1, mapS1
// build map of occurence in s2 with the size of len s1
// slide from left to right in s2
//  loop in mapS1, check if exist in s2

// TC O(n*26)

func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) { return false }
    
    mapS1 := make(map[byte]int)
    mapS2 := make(map[byte]int)
    
    for i:=0;i<len(s1);i++ {
        mapS1[s1[i]] += 1
        mapS2[s2[i]] += 1
    }
    // fmt.Printf("%v\n%v\n\n", mapS1, mapS2)
    if checkMatch(mapS1, mapS2) { return true }
    
    for i:=len(s1);i<len(s2);i++ {
        mapS2[s2[i-len(s1)]] -= 1
        mapS2[s2[i]] += 1
        
        // fmt.Printf("%v\n%v\n\n", mapS1, mapS2)
        if checkMatch(mapS1, mapS2) { return true }
    }
    
    return false
}

func checkMatch(mapS1, mapS2 map[byte]int) bool {
    for k, v := range mapS1 {
        if mapS2[k] != v {
            return false
        }
    }
    return true
}

// -------

