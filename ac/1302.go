package ac

func deepestLeavesSum(root *TreeNode) int {
    if root == nil {
        return 0
    }

    d := []*TreeNode{root}
    r := 0

    for len(d) != 0 {
        n := len(d)
        r = 0
        for i := 0; i < n; i++ {
            node := d[0]
            d = d[1:]
            r += node.Val
            if node.Left != nil {
                d = append(d, node.Left)
            }
            if node.Right != nil {
                d = append(d, node.Right)
            }
        }

    }
    return r
}
