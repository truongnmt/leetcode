// bfs:       1 2 3 N N 4 5 N N N N
// pre order: 1 2 N N 3 4 N N 5 N N
//            ^

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    // preorder
    nodes := []string{}
    
    var dfs func(*TreeNode) 
    dfs = func(node *TreeNode) {
        if node == nil {
            nodes = append(nodes, "N")
            return
        }
        
        nodes = append(nodes, strconv.Itoa(node.Val))
        dfs(node.Left)
        dfs(node.Right)
    }
    
    dfs(root)
    
    return strings.Join(nodes, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {        
    var nodes []string
    nodes = strings.Split(data, ",")
    
    i := 0
    
    // 1 2 N N 3 4 N N 5 N N
    //     i
    var dfs func() *TreeNode
    dfs = func() *TreeNode {
        if i>=len(nodes) || nodes[i] == "N" {
            i++
            return nil
        }
        
        val, _ := strconv.Atoi(nodes[i])
        i++
        node := &TreeNode {
            Val: val,
            Left: dfs(),
            Right: dfs(),
        }
        
        return node
    }
    
    return dfs()
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */