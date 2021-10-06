func numIslands(grid [][]byte) int {
    island := 0
    m, n := len(grid), len(grid[0])
    drs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    visited := make(map[[2]int]bool)

    var dfs func(int, int)
    dfs = func(i, j int) {
        visited[[2]int{i,j}] = true
        for _, dr := range(drs) {
            x, y := i+dr[0], j+dr[1]
            if 0<=x && x<m && 0<=y && y<n && string(byte(grid[x][y]))=="1" {
                if visited[[2]int{x,y}] == false {
                    dfs(x, y)
                }
            }
        }
    }

    for i:=0;i<m;i++ {
        for j:=0;j<n;j++ {
            if string(byte(grid[i][j]))=="1" && visited[[2]int{i,j}] == false {
                dfs(i, j)
                island += 1
            }
        }
    }

    return island
}
