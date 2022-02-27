// hit cog
// hot dot dog lot log cog

// queue: (hot dot lot dog) log cog
// seen: hot dot lot dog
// change: 4


func ladderLength(beginWord string, endWord string, wordList []string) int {
    
    adjacents := make(map[string][]string)
    
    maskChar := func(s string, i int) string {
        return s[0:i]+"*"+s[i+1:]
    }
    
    endWordInWordList := false
    for _, w := range wordList {
        if w == endWord { endWordInWordList = true }
        for i:=0; i<len(w);i++ {
            maskedChar := maskChar(w, i)
            if list, exist := adjacents[maskedChar]; exist {
                adjacents[maskedChar] = append(list, w)
            } else {
                adjacents[maskedChar] = []string{w}
            }
        }
    }
    
    fmt.Println(adjacents)
    
    if !endWordInWordList || len(adjacents) == 0 {
        return 0
    }
    
    q := []string{beginWord}
    seen := make(map[string]bool)
    change := 0
    
    for len(q) > 0 {
        change += 1
        tmp := []string{}
        for _, w := range q {
            if w == endWord { return change }
            // fmt.Printf("w=%v seen[w]=%v\n", w, seen[w])
            if !seen[w] {
                for i:=0;i<len(w);i++ {
                    maskedChar := maskChar(w, i)
                    tmp = append(tmp, adjacents[maskedChar]...)
                }
            }
            seen[w] = true
        }
        q = tmp
        // fmt.Printf("q=%v seen=%v change=%v\n", q, seen, change)
    }
    
    return 0
}




