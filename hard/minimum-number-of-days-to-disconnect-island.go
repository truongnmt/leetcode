func minDays(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    directions := [][]int{{-1,0}, {1,0}, {0,-1}, {0,1}}
    
    var dfs func(int, int, map[[2]int]bool)
    dfs = func(i, j int, visited map[[2]int]bool) {
        visited[[2]int{i,j}] = true
        for _, dr := range(directions) {
            x, y := i+dr[0], j+dr[1]
            if 0<=x && x<m && 0<=y && y<n && grid[x][y]==1 && visited[[2]int{x,y}]==false {
                dfs(x,y, visited)
            }
        }
    }
    
    var islandCount func() int
    islandCount = func() (int) {
        island := 0
        visited := make(map[[2]int]bool)
        for i:=0;i<m;i++ {
            for j:=0;j<n;j++ {
                if grid[i][j]==1 && visited[[2]int{i,j}]==false {
                    dfs(i,j, visited)
                    island+=1
                }
            }
        }
        if len(visited) == 0 {
            return 0
        }
        return island
    }
    
    currentIsland := islandCount()
    
    if currentIsland > 1 {
        return 0
    }
    
    for i:=0;i<m;i++ {
        for j:=0;j<n;j++ {
            if grid[i][j]==1 {
                grid[i][j]=0
                newIslandCount := islandCount()
                if newIslandCount > 1 || newIslandCount == 0 {
                    return 1
                }
                grid[i][j]=1
            }
        }
    }
    
    return 2
}