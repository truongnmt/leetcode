/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    
    var res [][]int
    st := []*TreeNode{root}
    timer := 0
    for len(st)>0 {
        timer += 1
        levelNode := []int{}
        if timer % 2 != 0 {
            for _, node := range(st) {
                levelNode = append(levelNode, node.Val)
            }
        } else {
            for i:=len(st)-1;i>=0;i-- {
                levelNode = append(levelNode, st[i].Val)
            }
        }
        res = append(res, levelNode)
        
        tmp := []*TreeNode{}
        for _, node := range(st) {
            if node.Left != nil {
                tmp = append(tmp, node.Left)
            }
            if node.Right != nil {
                tmp = append(tmp, node.Right)
            }
        }
        st = tmp
    }
    
    return res
}
