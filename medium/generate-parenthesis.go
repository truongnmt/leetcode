func generateParenthesis(n int) []string {
    ans := []string{}
    
    var generate func(int, int, string)
    generate = func(openCount int, closeCount int, input string) {
        if openCount == closeCount && openCount == n {
            ans = append(ans, input)
            return
        }
        
        if openCount < n {
            generate(openCount+1, closeCount, input+"(")
        }
        if closeCount < openCount {
            generate(openCount, closeCount+1, input+")")
        }
    }
    generate(1, 0, "(")
    
    return ans
}
