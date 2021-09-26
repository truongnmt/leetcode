/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderTraversal(root *TreeNode) []int {
    nodes := dfs(root)
    return nodes
}

func dfs(u *TreeNode) []int {
    if u == nil {
        return nil
    }
    
    l := dfs(u.Left)
    nodes := []int{u.Val}
    nodes = append(l, nodes...)
    r := dfs(u.Right)
    nodes = append(nodes, r...)
    return nodes
}

// iterative

func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    
    s := []*TreeNode{root}
    nodes := []int{}
    seen := make(map[*TreeNode]bool)
    
    for len(s) > 0 {
        cur := s[len(s)-1]
        s = s[:(len(s)-1)]
        if cur.Left == nil || seen[cur] == true {
            nodes = append(nodes, cur.Val)
            if cur.Right != nil {
                s = append(s, cur.Right)
            }
        } else {
            s = append(s, cur)
            s = append(s, cur.Left)
            seen[cur] = true
        }
    }
    
    return nodes
}

