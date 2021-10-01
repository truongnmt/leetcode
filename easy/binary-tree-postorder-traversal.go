/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func postorderTraversal(root *TreeNode) []int {
    nodes := dfs(root)
    return nodes
}

func dfs(u *TreeNode) []int {
    if u == nil {
        return nil
    }
    
    lnodes := dfs(u.Left)
    rnodes := dfs(u.Right)
    nodes := append(lnodes, rnodes...)
    nodes = append(nodes, u.Val)
    return nodes
}

// iterative

func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    
    s := []*TreeNode{root}
    nodes := []int{}
    
    for len(s) > 0 {
        cur := s[len(s)-1]
        s = s[:len(s)-1]
        nodes = append([]int{cur.Val}, nodes...)
        if cur.Left != nil {
            s = append(s, cur.Left)
        }
        if cur.Right != nil {
            s = append(s, cur.Right)
        }
    }
    
    return nodes
}

