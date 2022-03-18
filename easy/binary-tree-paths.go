/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func binaryTreePaths(root *TreeNode) []string {
    
    ans := []string{}
    
    var dfs func(*TreeNode, string)
    dfs = func(node *TreeNode, currentPath string) {
        if node.Left == nil && node.Right == nil {
            currentPath += strconv.Itoa(node.Val)
            ans = append(ans, currentPath)
            return
        }
        
        if node.Left != nil {
            dfs(node.Left, fmt.Sprintf("%v%v->", currentPath, strconv.Itoa(node.Val)))
        }
        if node.Right != nil {
            dfs(node.Right, fmt.Sprintf("%v%v->", currentPath, strconv.Itoa(node.Val)))
        }
    }
    
    dfs(root, "")
    
    return ans
}

// ==========================================================

func binaryTreePaths(root *TreeNode) []string {
    
    ans := []string{}
    
    var dfs func(*TreeNode, []string)
    dfs = func(node *TreeNode, currentPath []string) {
        currentPath = append(currentPath, strconv.Itoa(node.Val))
        
        if node.Left == nil && node.Right == nil {
            ans = append(ans, strings.Join(currentPath, ""))
            return
        }
        
        if node.Left != nil {
            dfs(node.Left, append(currentPath, "->"))
        }
        if node.Right != nil {
            dfs(node.Right, append(currentPath, "->"))
        }
    }
    
    dfs(root, []string{})
    
    return ans
}

// ==========================================================

func binaryTreePaths(root *TreeNode) []string {
    
    ans := []string{}
    
    var dfs func(*TreeNode, []string)
    dfs = func(node *TreeNode, currentPath []string) {
        currentPath = append(currentPath, strconv.Itoa(node.Val))
        
        if node.Left == nil && node.Right == nil {
            ans = append(ans, strings.Join(currentPath, ""))
            return
        }
        
        if node.Left != nil {
            dfs(node.Left, append(currentPath, "->"))
        }
        if node.Right != nil {
            dfs(node.Right, append(currentPath, "->"))
        }
    }
    
    dfs(root, []string{})
    
    return ans
}
