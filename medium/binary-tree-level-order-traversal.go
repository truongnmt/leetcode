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

// bfs in queue

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    var res [][]int
    queue := []*TreeNode{root}
    for level:=0;len(queue)>0;level++{
        res = append(res, []int{})
        for levelSize:=len(queue);levelSize>0;levelSize--{
            if queue[0].Left != nil {
                queue = append(queue, queue[0].Left)
            }
            if queue[0].Right != nil {
                queue = append(queue, queue[0].Right)
            }
            res[level] = append(res[level], queue[0].Val)
            queue = queue[1:]
        }
    }
    return res
}

// bfs in stack

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    var res [][]int
    st := []*TreeNode{root}
    for len(st)>0 {
        levelItem := []int{}
        for _, node := range(st) {
            levelItem = append(levelItem, node.Val)
        }
        res = append(res, levelItem)
        
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
