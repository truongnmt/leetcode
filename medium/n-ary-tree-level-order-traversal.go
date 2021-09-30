/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 func levelOrder(root *Node) [][]int {
    if root == nil {
        return nil
    }
    
    var res [][]int
    st := []*Node{root}
    for len(st)>0 {
        nodeLevel := []int{}
        tmp := []*Node{}
        for _, node := range(st) {
            nodeLevel = append(nodeLevel, node.Val)
            tmp = append(tmp, node.Children...)
        }
        res = append(res, nodeLevel)
        st = tmp
    }
    return res
}