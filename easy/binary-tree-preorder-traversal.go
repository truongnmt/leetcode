/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func preorderTraversal(root *TreeNode) []int {
    nodes := dfs(root)
    return nodes
}

func dfs(u *TreeNode) []int {
    if u == nil {
        return nil
    }
    
    nodes := []int{u.Val}
    lnodes := dfs(u.Left)
    rnodes := dfs(u.Right)
    nodes = append(nodes, lnodes...)
    nodes = append(nodes, rnodes...)
    return nodes
}

// iterative

func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    
    s := []*TreeNode{root}
    nodes := []int{}
    
    for len(s) > 0 {
        cur := s[len(s)-1]
        s = s[:len(s)-1]
        nodes = append(nodes, cur.Val)
        if cur.Right != nil {
            s = append(s, cur.Right)
        }
        if cur.Left != nil {
            s = append(s, cur.Left)
        }
    }
    
    return nodes
}

