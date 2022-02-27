// A rectangular map consisting of N rows and M columns of square areas is given. Each area is painted with a certain color. 

// Two areas on the map belong to the same country if the following conditions are met: 
// - they have the same color
// - it is possible to travel from one area to the other by moving only north, south, west or east without moving over areas of a different color.

// The map can be described by a zero-indexed matrix A consisting of N rows and M columns of integers. The color of each area is described by the corresponding element of the matrix. Two areas have the same color if and only if their corresponding matrix elements have the same value. 

// For example, consider the following matrix A consisting of seven rows and three columns: 

// A[0][0] = 5 A[0][1] = 4 A[0][2] = 4 
// A[1][0] = 4 A[1][1] = 3 A[1][2] = 4 
// A[2][0] = 3 A[2][1] = 2 A[2][2] = 4 
// A[3][0] = 2 A[3][1] = 2 A[3][2] = 2 
// A[4][0] = 3 A[4][1] = 3 A[4][2] = 4 
// A[5][0] = 1 A[5][1] = 4 A[5][2] = 4 
// A[6][0] = 4 A[6][1] = 1 A[6][2] = 1 

// Matrix A describes a map that is colored with five colors. The areas on the map belong to eleven different countries (C1−C11), as shown in the following figure: 

// area  A[0][0] forms a one-area country;
// areas A[0][1], A[0][2], A[1][2], A[2][2] belong to the same country;
// area  A[1][0] forms a one-area country;
// area  A[1][1] forms a one-area country;
// area  A[2][0] forms a one-area country;
// areas A[2][1], A[3][0], A[3][1], A[3][2] belong to the same country;
// areas A[4][0], A[4][1] belong to the same country;
// areas A[4][2], A[5][1], A[5][2] belong to the same country;
// area  A[5][0] forms a one-area country;
// area  A[6][0] forms a one-area country;
// areas A[6][1], A[6][2] belong to the same country.

// Write a function that, given a zero-indexed matrix A consisting of N rows and M columns of integers, returns the number of different countries to which the areas of the map described by matrix A belong. 

// N is an integer within the range [1..1,000,000];
// M is an integer within the range [1..1,000,000];
// the number of elements in matrix A is within the range [1..1,000,000];
// each element of matrix A is an integer within the range
//  [-1,000,000,000..1,000,000,000].

// For example, given matrix A consisting of seven rows and three columns corresponding to the example above, the function should return 11

// Complexity:
// expected worst-case time complexity is O(N*M);
// expected worst-case space complexity is O(N*M).

// 5 4 4
// 4 3 4
// 3 2 4
// 2 2 2
// 3 3 4
// 1 4 4
// 4 1 1

// nums = [][]int
// mapSeen[[2]int]bool
// numCountry
// directions = [4][2]int{[-1,0],[0,1],[1,0],[0,-1]}

// func dfs(x, y int) {
//     for direction in directions
//     neighbor = x+direction[0], y+direction[1]
//     if neighbor == countryCode {
//         add neighbor to mapSeen
//         dfs(neighbor)
//     }
// }

// loop in row
// loop in col
// if not in mapSeen
//  add to mapSeen
//  dfs(x, y)
//  num++

func main() {
    nums := [][]int{{5,4,4},{4,3,4},{3,2,4},{2,2,2},{3,3,4},{1,4,4},{4,1,1}}
    fmt.Println(numberOfCountry(nums))
}

func numberOfCountry(nums [][]int) int {
    mapSeen := make(map[[2]int]bool)
    numCountry := 0
    directions := [4][2]int{{-1,0},{0,1},{1,0},{0,-1}}
    
    var dfs func(int, int)
    dfs = func(x, y int) {
        fmt.Printf("x=%v y=%v\n", x,y)
        for _, direction := range directions {
            neighborCordinateX := x+direction[0]
            neighborCordinateY := y+direction[1]
            _, seen := mapSeen[[2]int{neighborCordinateX, neighborCordinateY}]
            if 0<=neighborCordinateX && neighborCordinateX < len(nums) &&
              0<=neighborCordinateY && neighborCordinateY < len(nums[0]) &&
              nums[x][y] == nums[neighborCordinateX][neighborCordinateY] && !seen{
                  mapSeen[[2]int{neighborCordinateX, neighborCordinateY}] = true
                  dfs(neighborCordinateX, neighborCordinateY)
            }
        }
    }
    
    for x, col := range nums {
        for y, _ := range col {
            if _, seen := mapSeen[[2]int{x,y}]; !seen {
                mapSeen[[2]int{x,y}] = true
                dfs(x, y)
                numCountry++
            }
        }
    }
    
    return numCountry
}
