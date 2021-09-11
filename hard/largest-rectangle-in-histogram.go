// Time Limit Exceeded: O(n^2)
func largestRectangleArea1(heights []int) int {
    if len(heights)==1 {
        return heights[0]
    }
    
    maxArea:=0
    for i:=0;i<len(heights);i++ {
        minHeight := heights[i]
        for j:=i;j<len(heights);j++{
            if heights[j] < minHeight {
                minHeight=heights[j]
            }
            area := minHeight*(j-i+1)
            if area > maxArea {
                maxArea = area
            }
        }
    }
    return maxArea
}

// O(n)
func largestRectangleArea(heights []int) int {
    minStack := []int{}
    a:=heights
    leftSmaller := make([]int, len(a))
    rightSmaller := make([]int, len(a))
    
    // build leftSmaller
    for i:=0;i<len(a);i++ {
        // fmt.Printf("i:%v a[i]:%v minStack:%v\n", i, a[i], minStack)
        if len(minStack)==0 {
            minStack=append(minStack, i)
            leftSmaller[i]= 0
            continue
        } else if a[i] > a[minStack[len(minStack)-1]] {
            minStack=append(minStack, i)
            leftSmaller[i]= i
            continue
        }
        
        for len(minStack)>0 && a[i] <= a[minStack[len(minStack)-1]] {
            if len(minStack)==1 {
                leftSmaller[i]=0
            } else {
                leftSmaller[i]=leftSmaller[minStack[(len(minStack)-1)]]
            }
            minStack=minStack[:(len(minStack)-1)]
        }
        minStack=append(minStack, i)
        if len(minStack)==0 {
            leftSmaller[i]= 0
            continue
        }
    }
    // fmt.Println(leftSmaller)
    
    minStack=[]int{}
    // build rightSmaller
    for i:=len(a)-1;i>=0;i-- {
        // fmt.Printf("i:%v a[i]:%v minStack:%v\n", i, a[i], minStack)
        if len(minStack)==0 {
            minStack=append(minStack, i)
            rightSmaller[i]=len(a)-1
            continue
        } else if a[i] > a[minStack[len(minStack)-1]] {
            minStack=append(minStack, i)
            rightSmaller[i]=i
            continue
        }
        
        for len(minStack)>0 && a[i] <= a[minStack[len(minStack)-1]] {
            if len(minStack)==1 {
                rightSmaller[i]=len(a)-1
            } else {
                rightSmaller[i]=rightSmaller[minStack[(len(minStack)-1)]]
            }
            minStack=minStack[:(len(minStack)-1)]
        }
        minStack=append(minStack, i)
        if len(minStack)==0 {
            rightSmaller[i]=len(a)-1
            continue
        }
    }
    // fmt.Println(rightSmaller)
    
    maxArea:=0
    for i:=0;i<len(leftSmaller);i++ {
        area:=(rightSmaller[i]-leftSmaller[i]+1)*a[i]
        // fmt.Printf("area:%v maxArea:%v\n", area, maxArea)
        if area > maxArea {
            maxArea=area
        }
    }
    
    return maxArea
}

// ---------------------------------------------------------------------------------
// O(N) using go channel
func buildLeftSmaller(heights []int, c chan []int) {
    a := heights
    n := len(heights)
    leftSmaller := make([]int, n)
    leftSmaller[0] = 0
    minStack := []int{0}
    for i:=1;i<n;i++ {
        if len(minStack)==0 {
            minStack=append(minStack, i)
            leftSmaller[i]= 0
            continue
        } else if a[i] > a[minStack[len(minStack)-1]] {
            minStack=append(minStack, i)
            leftSmaller[i]= i
            continue
        }
        
        for len(minStack)>0 && a[i] <= a[minStack[len(minStack)-1]] {
            if len(minStack)==1 {
                leftSmaller[i]=0
            } else {
                leftSmaller[i]=leftSmaller[minStack[(len(minStack)-1)]]
            }
            minStack=minStack[:(len(minStack)-1)]
        }
        minStack=append(minStack, i)
        if len(minStack)==0 {
            leftSmaller[i]= 0
            continue
        }
    }
    c <- leftSmaller
}

func buildRightSmaller(heights []int, c chan []int) {
    a := heights
    n := len(heights)
    rightSmaller := make([]int, n)
    rightSmaller[n-1] = n-1
    minStack := []int{n-1}
    for i:=n-2;i>=0;i--{
        if len(minStack)==0 {
            minStack=append(minStack, i)
            rightSmaller[i]=len(a)-1
            continue
        } else if a[i] > a[minStack[len(minStack)-1]] {
            minStack=append(minStack, i)
            rightSmaller[i]=i
            continue
        }
        
        for len(minStack)>0 && a[i] <= a[minStack[len(minStack)-1]] {
            if len(minStack)==1 {
                rightSmaller[i]=len(a)-1
            } else {
                rightSmaller[i]=rightSmaller[minStack[(len(minStack)-1)]]
            }
            minStack=minStack[:(len(minStack)-1)]
        }
        minStack=append(minStack, i)
        if len(minStack)==0 {
            rightSmaller[i]=len(a)-1
            continue
        }
    }
    c <- rightSmaller
}

func largestRectangleArea(heights []int) int {
    c1 := make(chan []int)
    c2 := make(chan []int)
    
    go buildLeftSmaller(heights, c1)
    leftSmaller := <- c1
    go buildRightSmaller(heights, c2)
    rightSmaller := <- c2

    maxArea := 0
    for i:=0;i<len(heights);i++ {
        area := (rightSmaller[i] - leftSmaller[i] + 1) * heights[i]
        if area > maxArea {
            maxArea = area
        }
    }

    return maxArea
}
