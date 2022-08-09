package algo

func maxDepth(root *TreeNode) int {
    return maxd(root)
}

func maxd(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return 1
    }
    ml := maxd(root.Left)
    mr := maxd(root.Right)
    if ml > mr {
        return ml + 1
    } else {
        return mr + 1
    }
}
