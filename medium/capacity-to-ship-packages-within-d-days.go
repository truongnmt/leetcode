func calculateShipDays(cap int, weights []int) int {
    days := 0
    currentCap := 0
    for _, weight := range weights {
        if weight > cap {
            return 25000000
        }
        if currentCap + weight > cap {
            currentCap = 0
            days++
        }
        currentCap += weight
        // fmt.Printf("days=%v currentCap=%v weight=%v cap=%v\n", days, currentCap, weight, cap)
    }
    if currentCap > 0 {
        days++
    }
    
    return days
}

func shipWithinDays(weights []int, days int) int {
    
    // ship capacity required ranged from: 1 ... weights.length*weights[i] = 5*10^4*500=25000000
    // binary search from 1 to 25000000
    
    left, right := 0, 25000000
    for left+1 < right {
        pivot := (left+right)/2
        // fmt.Printf("left=%v pivot=%v right=%v cal=%v compare=%v\n", left, pivot, right, calculateShipDays(pivot, weights), calculateShipDays(pivot, weights) <= days)
        if calculateShipDays(pivot, weights) <= days {
            right = pivot
        } else {
            left = pivot
        }
    }
    
    return right
}
