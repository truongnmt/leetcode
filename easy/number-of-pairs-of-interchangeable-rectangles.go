func interchangeableRectangles(rectangles [][]int) int64 {
    m := make(map[float64]int)
    for _, r := range rectangles {
        m[float64(r[0])/float64(r[1])] += 1
    }
    
    var result int64 = 0
    for _, v := range m {
        if v > 1 {
            result += int64(v*(v-1)/2)
        }
    }
    return int64(result)
}
