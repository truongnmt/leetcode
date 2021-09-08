func binaryGen(x, n int, a, nums []int) int {
	if x == n {
        sum := 0
        for i:=0;i<len(a);i++ {
            if a[i] == 1 {
                sum ^= nums[i]
            }
        }
        // fmt.Printf("a: %v, sum: %v\n", a, sum)
		return sum
	}
    
    sum := 0
	for i:=0;i<2;i++ {
		a[x] = i
        sum += binaryGen(x+1, n, a, nums)
	}
    
    return sum
}


func subsetXORSum(nums []int) int {
    n := len(nums)
	a := make([]int, n)
    sum := binaryGen(0, n, a, nums)
    return sum
}
