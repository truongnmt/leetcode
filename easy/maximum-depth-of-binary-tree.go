/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxDepth(root *TreeNode) int {
    max := dfs(root, 0)
    return max
}

func dfs(u *TreeNode, depth int) int {
    if u == nil {
        return depth
    }
    
    l := dfs(u.Left, depth+1)
    r := dfs(u.Right, depth+1)
    if l > r {
        return l
    } else {
        return r
    }
    
    return 0
}
