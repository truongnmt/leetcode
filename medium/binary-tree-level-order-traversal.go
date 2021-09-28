/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    
    m := make(map[int][]int)
    dfs(root, 1, m)
    
    keys := make([]int, 0)
    ans := make([][]int, 0)
    for k := range m {
        keys = append(keys, k)
    }
    sort.Ints(keys)
    
    for _, k := range keys {
        ans = append(ans, m[k])
    }
    
    return ans
}

func dfs(u *TreeNode, h int, m map[int][]int) {
    if u == nil {
        return
    }
    
    m[h] = append(m[h], u.Val)
    dfs(u.Left, h+1, m)
    dfs(u.Right, h+1, m)
}
