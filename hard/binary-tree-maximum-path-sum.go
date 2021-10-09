/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxPathSum(root *TreeNode) int {
    if root == nil {return 0}
    
    max := root.Val
    
    var getMaxPath func(*TreeNode) int
    getMaxPath = func(u *TreeNode) int {
        if u == nil {
            return 0
        }
        
        l := getMaxPath(u.Left)
        r := getMaxPath(u.Right)
        if l<0 {l=0}
        if r<0 {r=0}
        nodeMax := l + u.Val + r
        // fmt.Printf("l=%v val=%v r=%v nodeMax=%v max=%v\n", l, u.Val, r, nodeMax, max)
        if max < nodeMax { max = nodeMax }
        
        if l > r {
            return u.Val+l
        }
        
        return u.Val+r
    }
    
    getMaxPath(root)
    return max
}