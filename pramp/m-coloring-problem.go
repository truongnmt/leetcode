// Given an undirected graph and a number m, determine if the graph can be coloured with at most m colours such that no two adjacent vertices of the graph are colored with the same color. Here coloring of a graph means the assignment of colors to all vertices. 

// Input-Output format: 

// Input: 

// A 2D array graph[V][V] where V is the number of vertices in graph and graph[V][V] is an adjacency matrix representation of the graph. A value graph[i][j] is 1 if there is a direct edge from i to j, otherwise graph[i][j] is 0.
// An integer m is the maximum number of colors that can be used.
// Output: 
// An array color[V] that should have numbers from 1 to m. color[i] should represent the color assigned to the ith vertex. The code should also return false if the graph cannot be colored with m colors.

// Input:  
// graph = {0, 1, 1, 1},
//         {1, 0, 1, 0},
//         {1, 1, 0, 1},
//         {1, 0, 1, 0}
// Output: 
// Solution Exists: 
// Following are the assigned colors
//  1  2  3  2

// ---

// create a colorVertext func
// assign a vertex a color
// if its valid (adj vertex color is different)
//   recursively assign adj vertex color
// else try another color -> from color 1 -> m
// base case is if vertex == size return true


func main() {
    graph := [][]int{
        { 0, 1, 1, 1 },
        { 1, 0, 1, 0 },
        { 1, 1, 0, 1 },
        { 1, 0, 1, 0 },
    };
    m := 3
    fmt.Println(colorVertexCheck(graph, m))
}

func colorVertexCheck(graph [][]int, m int) bool {
    numberOfVertex := len(graph)
    // vertexColor := [m]int{} -> not work Line 45: Char 20: non-constant array bound m (solution.go)
    vertexColor := make([]int, m+1)
    
    var coloringVertex func(int) bool
    coloringVertex = func(vertex int) bool {
        fmt.Printf("\nvertex=%v numberOfVertex=%v\n", vertex, numberOfVertex)
        if vertex == numberOfVertex {
            return true
        }
        
        // try color i for this vertex
        for color:=1;color<=m;color++ {
            // check if color valid
            valid := true
            for vertexDestination, connection := range graph[vertex] {
                // if have connection and connected vertex's color is the same color -> not valid
                fmt.Printf("connection=%v color=%v vertexDestination=%v vertexColor=%v\n", connection, color, vertexDestination, vertexColor)
                if connection == 1 && color == vertexColor[vertexDestination] { 
                    valid = false
                    break
                }
            }
            
            if valid {
                vertexColor[vertex] = color
                fmt.Printf("vertexColor=%v\n", vertexColor)
                return coloringVertex(vertex+1)
            }
        }
        
        return false
    }
    
    return coloringVertex(0)
}


// Method 3:  Using BFS

// The approach here is to color each node from 1 to n initially by color 1. And start travelling BFS from an unvisited starting node to cover all connected components in one go. On reaching each node during BFS traversal, do the following:

// Check all edges of the given node.
// For each vertex connected to our node via an edge:
// check if the color of the nodes is the same. If same, increase the color of the other node (not the current) by one.
// check if it visited or unvisited. If not visited, mark it as visited and push it in a queue.
// Check condition for maxColors till now. If it exceeds M, return false

