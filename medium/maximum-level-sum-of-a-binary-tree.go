/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxLevelSum(root *TreeNode) int {
    max := root.Val
    st := []*TreeNode{root}
    level := 1
    timer := 0
    for len(st)>0 {
        var sumNodeLevel int
        timer += 1
        tmp := []*TreeNode{}
        for _, node := range(st) {
            sumNodeLevel += node.Val
            if node.Left != nil { tmp = append(tmp, node.Left) }
            if node.Right != nil { tmp = append(tmp, node.Right) }
        }
        st = tmp
        
        if max < sumNodeLevel {
            max = sumNodeLevel
            level = timer
        }
    }
    
    return level
}
