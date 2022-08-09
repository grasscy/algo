package algo

func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    zhongxu(root, &res)
    return res
}

func zhongxu(root *TreeNode, res *[]int) {
    if root == nil {
        return
    }
    zhongxu(root.Left, res)
    *res = append(*res, root.Val)
    zhongxu(root.Right, res)
}
