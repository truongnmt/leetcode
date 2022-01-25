func longestPalindrome(s string) string {
    // loop in each char in s
    // set left and right index, move to left and right while left = right
    // set left and right index for i+1, move to left and right while left = right
    // get max index
    
    ans := [2]int{}
    for i, _ := range s {
        left := i-1
        right := i+1
        for 0<=left && left < right && right < len(s) {
            // fmt.Printf("i=%v left=%v right=%v\n", i, left, right)
            if s[left] != s[right] {
                break
            } else {
                if right-left > ans[1]-ans[0] {
                    ans = [2]int{left, right}
                }
            }
            left--
            right++
        }
        
        if i+1 < len(s) {
            left := i
            right := i+1
            for 0<=left && left < right && right < len(s) {
                // fmt.Printf("i=%v left=%v right=%v\n", i, left, right)
                if s[left] != s[right] {
                    break
                } else {
                    if right-left > ans[1]-ans[0] {
                        ans = [2]int{left, right}
                    }
                }
                left--
                right++
            }
        }
    }
    
    // fmt.Println(ans)
    
    return s[ans[0]:ans[1]+1]
}
