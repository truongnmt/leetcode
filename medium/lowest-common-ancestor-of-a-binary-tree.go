/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var ans *TreeNode
    
    var dfs func(*TreeNode) (bool, bool)
    dfs = func(u *TreeNode) (bool, bool) {
        // return bool, bool whether u is ancestor of p q
        if u==nil || ans != nil {
            return false, false
        }
        
        left_p, left_q := dfs(u.Left)
        right_p, right_q := dfs(u.Right)
        is_ancestor_of_p, is_ancestor_of_q := false, false
        if left_p || right_p || u==p {
            is_ancestor_of_p = true
        }
        if left_q || right_q || u==q {
            is_ancestor_of_q = true
        }
        if is_ancestor_of_p && is_ancestor_of_q && ans == nil {
            ans = u
        }
        return is_ancestor_of_p, is_ancestor_of_q
    }
    
    _, _ = dfs(root)
    return ans
}