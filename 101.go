package algo

func isSymmetric(root *TreeNode) bool {
    return duicheng(root.Left, root.Right)
}

func duicheng(l *TreeNode, r *TreeNode) bool {
    if l == nil && r == nil {
        return true
    }
    if l == nil || r == nil {
        return false
    }
    return l.Val == r.Val && duicheng(l.Left, r.Right) && duicheng(l.Right, r.Left)
}
